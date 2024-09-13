CREATE TYPE email_event_module AS ENUM ('interview', 'application', 'job_request');

CREATE TYPE email_event_action AS ENUM (
  'create', 'update', 'cancel',
  'cd_applied', 'cd_interviewing', 'cd_offering', 'cd_failed_cv', 'cd_failed_interview', 'cd_offer_lost', 'cd_hired',
  'close', 'open', 'reopen', 'need_approval'
);

CREATE TABLE IF NOT EXISTS email_events (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  module email_event_module,
  action email_event_action,
  name TEXT,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP
);

INSERT INTO email_events (id, module, action, name) VALUES
  ('2e25bedd-a87d-497c-aa39-1e9ee5879642', 'interview', 'create', 'model.events.interview.create'),
  ('466f2006-d51e-4a16-98cf-6ff9b73838c7', 'interview', 'update', 'model.events.interview.update'),
  ('ae4055be-72b2-46f3-9de6-6ee9789cc6b2', 'interview', 'cancel', 'model.events.interview.cancel'),
  ('1f1120bd-32af-4a74-ae96-e544efdce755', 'application', 'cd_applied', 'model.events.application.cd_applied'),
  ('d43349e1-8062-4b97-99de-1fbfa23183dd', 'application', 'cd_interviewing', 'model.events.application.cd_interviewing'),
  ('097520f8-c92c-4753-9562-e40e83d1de5a', 'application', 'cd_offering', 'model.events.application.cd_offering'),
  ('1b5f4520-638f-4225-a1e7-e5b8db09b1be', 'application', 'cd_failed_cv', 'model.events.application.cd_failed_cv'),
  ('d5a6ded0-6130-4b31-a950-70cf67819230', 'application', 'cd_failed_interview', 'model.events.application.cd_failed_interview'),
  ('d30339a5-8441-4b90-b5af-4c3c69189346', 'application', 'cd_offer_lost', 'model.events.application.cd_offer_lost'),
  ('c989f71e-0507-441f-991f-f0b20a61605a', 'application', 'cd_hired', 'model.events.application.cd_hired'),
  ('746c015e-8c09-4e02-a39e-046140abd9e2', 'job_request', 'create', 'model.events.job_req.create'),
  ('834f7a93-e7b7-461c-bbcc-1eb69246322b', 'job_request', 'update', 'model.events.job_req.update'),
  ('940c79c7-9cab-4478-9ba0-d67c4dd8edc8', 'job_request', 'close', 'model.events.job_req.close'),
  ('bb085343-d9ee-4358-a3af-8a51359be03c', 'job_request', 'open', 'model.events.job_req.open'),
  ('806899b6-957f-4741-8e8c-fd6b00abe04d', 'job_request', 'reopen', 'model.events.job_req.reopen'),
  ('96946d24-5b02-4f3d-8664-3caaed7cc587', 'job_request', 'need_approval', 'model.events.job_req.need_approval');
