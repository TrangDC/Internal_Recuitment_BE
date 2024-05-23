-- Create the failed_reason_enum type
CREATE TYPE failed_reason_enum AS ENUM ('poor_professionalism','poor_fit_and_engagement','over_expectations','over_qualification','language_deficiency','weak_technical_skills','poor_interpersonal_skills','poor_problem_solving_skills','poor_management_skills','candidate_withdrawal','others');

-- Add the failed_reason field to the candidate_jobs table
ALTER TABLE candidate_jobs ADD COLUMN failed_reason jsonb NULL;

CREATE OR REPLACE FUNCTION validate_failed_reason()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.failed_reason IS NOT NULL THEN
        PERFORM 1
        FROM jsonb_array_elements_text(NEW.failed_reason) AS elem
        WHERE elem::failed_reason_enum IS NULL;

        IF FOUND THEN
            RAISE EXCEPTION 'Invalid value in failed_reason';
        END IF;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER validate_failed_reason_trigger
BEFORE INSERT OR UPDATE ON candidate_jobs
FOR EACH ROW EXECUTE FUNCTION validate_failed_reason();