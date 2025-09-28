CREATE TABLE IF NOT EXISTS posts (
  id uuid PRIMARY KEY,
  title text NOT NULL,
  user_id uuid NOT NULL,
  content text NOT NULL,
  created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
  updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);