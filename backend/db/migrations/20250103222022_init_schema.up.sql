CREATE TABLE "account" (
    "id" uuid UNIQUE PRIMARY KEY DEFAULT gen_random_uuid (),
    "email" varchar NOT NULL,
    "password" varchar NOT NULL,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE "player_card" (
    "id" serial PRIMARY KEY,
    "account_id" uuid NOT NULL,
    "session_id" uuid NOT NULL,
    "nickname" varchar NOT NULL,
    "preferences" varchar
);
CREATE TABLE "session" (
    "id" uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    "admin_id" uuid NOT NULL,
    "name" varchar NOT NULL DEFAULT 'Secret Santa game',
    "picture" varchar NOT NULL DEFAULT 'http://placehold.it/32x32',
    "seed" varchar,
    "is_archived" boolean NOT NULL DEFAULT 'false',
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP
);
ALTER TABLE "player_card"
ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");
ALTER TABLE "player_card"
ADD FOREIGN KEY ("session_id") REFERENCES "session" ("id");
ALTER TABLE "session"
ADD FOREIGN KEY ("admin_id") REFERENCES "account" ("id");