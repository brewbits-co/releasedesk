CREATE TABLE Releases
(
    ID             INTEGER PRIMARY KEY AUTOINCREMENT,
    ProductID      INTEGER
        CONSTRAINT Releases_Products_ID_fk REFERENCES Products,
    Version        TEXT     NOT NULL,
    TargetChannel  INTEGER
        CONSTRAINT Releases_Channel_ID_fk REFERENCES Channels,
    TargetPlatform TEXT,
    Status         TEXT     NOT NULL,
    ReleaseNotes   TEXT,
    CreatedAt      DATETIME NOT NULL,
    UpdatedAt      DATETIME NOT NULL
);

CREATE UNIQUE INDEX Releases_ProductID_Version_TargetPlatform_uindex
    ON Releases (ProductID, Version, TargetPlatform);