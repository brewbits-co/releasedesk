CREATE TABLE Releases
(
    ID            INTEGER PRIMARY KEY AUTOINCREMENT,
    AppID         INTEGER
        CONSTRAINT Releases_Apps_ID_fk REFERENCES app,
    Version       TEXT     NOT NULL,
    TargetChannel INTEGER
        CONSTRAINT Releases_Channel_ID_fk REFERENCES channel,
    Status        TEXT     NOT NULL,
    ReleaseNotes  TEXT,
    CreatedAt     DATETIME NOT NULL,
    UpdatedAt     DATETIME NOT NULL
);

CREATE UNIQUE INDEX Releases_AppID_Version_uindex
    ON Releases (AppID, Version);