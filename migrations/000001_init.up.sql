CREATE TABLE clients
(
    id         BIGSERIAL PRIMARY KEY,
    name       VARCHAR(100)                                       NOT NULL CHECK (length(name) >= 3),
    email      VARCHAR(255)                                       NOT NULL UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE agents
(
    id         BIGSERIAL PRIMARY KEY,
    name       VARCHAR(100)                                       NOT NULL CHECK (length(name) >= 3),
    email      VARCHAR(255)                                       NOT NULL UNIQUE,
    department VARCHAR(100)                                       NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE tickets
(
    id          BIGSERIAL PRIMARY KEY,
    client_id   BIGINT       NOT NULL REFERENCES clients (id) ON DELETE CASCADE,
    agent_id    BIGINT       REFERENCES agents (id) ON DELETE SET NULL,
    subject     VARCHAR(150) NOT NULL CHECK (length(subject) >= 5),
    description TEXT         NOT NULL CHECK (length(description) >= 10 AND length(description) <= 2000),
    status      VARCHAR(20)  NOT NULL    DEFAULT 'open' CHECK (status IN ('open', 'in_progress', 'resolved')),
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    resolved_at TIMESTAMP WITH TIME ZONE,

    CONSTRAINT check_resolved_at CHECK (
        (status = 'resolved' AND resolved_at IS NOT NULL AND resolved_at >= created_at) OR
        (status != 'resolved' AND resolved_at IS NULL)
        )
);

CREATE TABLE comments
(
    id          BIGSERIAL PRIMARY KEY,
    ticket_id   BIGINT                                             NOT NULL REFERENCES tickets (id) ON DELETE CASCADE,
    author_id   BIGINT                                             NOT NULL,
    author_type VARCHAR(20)                                        NOT NULL CHECK (author_type IN ('client', 'agent')),
    text        TEXT                                               NOT NULL,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_tickets_client_id ON tickets (client_id);
CREATE INDEX idx_tickets_agent_id ON tickets (agent_id);
CREATE INDEX idx_tickets_status ON tickets (status);
CREATE INDEX idx_comments_ticket_id ON comments (ticket_id);