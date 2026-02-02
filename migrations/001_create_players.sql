CREATE EXTENSION IF NOT EXISTS pgcrypto;


CREATE TABLE IF NOT EXISTS players (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  puuid TEXT NOT NULL UNIQUE,
  region TEXT NOT NULL CHECK (region IN ('europe','americas','asia')),
  game_name TEXT NOT NULL,
  tag_line TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS players_riotid_idx
  ON players (region, game_name, tag_line);