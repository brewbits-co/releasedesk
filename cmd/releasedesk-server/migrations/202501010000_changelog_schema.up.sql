CREATE TABLE Changelogs
(
    ID         INTEGER PRIMARY KEY AUTOINCREMENT,
    ReleaseID  INTEGER
        CONSTRAINT Changelogs_Releases_ID_fk REFERENCES Releases,
    Text       TEXT NOT NULL,
    ChangeType TEXT NOT NULL
);

CREATE INDEX Changelogs_ReleaseID_index
    ON Changelogs (ReleaseID);