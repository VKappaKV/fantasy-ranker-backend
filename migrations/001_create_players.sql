CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE players (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  puuid TEXT NOT NULL UNIQUE,
  region TEXT NOT NULL CHECK (region IN ('europe','americas','asia')),
  game_name TEXT NOT NULL,
  tag_line TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Per ricerche per RiotID
CREATE INDEX players_riotid_idx ON players (game_name, tag_line, region);
