CREATE TABLE LinkedBuilds
(
    ReleaseID INTEGER NOT NULL
        CONSTRAINT LinkedBuilds_Releases_ID_fk REFERENCES Releases ON DELETE CASCADE,
    BuildID   INTEGER NOT NULL
        CONSTRAINT LinkedBuilds_Builds_ID_fk REFERENCES Builds ON DELETE CASCADE,
    OS        TEXT    NOT NULL,
    PRIMARY KEY (ReleaseID, BuildID, OS)
);

CREATE INDEX LinkedBuilds_ReleaseID_index
    ON LinkedBuilds (ReleaseID);

CREATE INDEX LinkedBuilds_BuildID_index
    ON LinkedBuilds (BuildID);