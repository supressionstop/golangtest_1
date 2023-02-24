CREATE TABLE soccer
(
    created_at TIMESTAMP NOT NULL,
    provider VARCHAR NOT NULL,
    rate DECIMAL NOT NULL,
    PRIMARY KEY (created_at, provider, rate)
);
CREATE TABLE football
(
    created_at TIMESTAMP NOT NULL,
    provider VARCHAR NOT NULL,
    rate DECIMAL NOT NULL,
    PRIMARY KEY (created_at, provider, rate)
);
CREATE TABLE baseball
(
    created_at TIMESTAMP NOT NULL,
    provider VARCHAR NOT NULL,
    rate DECIMAL NOT NULL,
    PRIMARY KEY (created_at, provider, rate)
);
