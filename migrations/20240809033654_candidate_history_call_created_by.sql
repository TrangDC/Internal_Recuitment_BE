alter table candidate_history_calls add column created_by_id uuid not null references users(id);
