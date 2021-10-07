CREATE TABLE IF NOT EXISTS templates
(
    id uuid NOT NULL,
    "isWorkflow" boolean,
    name character varying,
    path character varying UNIQUE,
    contents text,
    "createdAt" timestamp with time zone,
    "updatedAt" timestamp with time zone,
    CONSTRAINT templates_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS targets
(
    id uuid NOT NULL,
    name character varying,
    "rawPath" character varying UNIQUE,
    "totalHosts" bigint,
    "createdAt" timestamp with time zone,
    "updatedAt" timestamp with time zone,
    CONSTRAINT targets_pkey PRIMARY KEY (id)
);