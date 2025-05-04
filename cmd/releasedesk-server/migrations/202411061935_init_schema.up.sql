CREATE TABLE Users
(
    ID        INTEGER PRIMARY KEY AUTOINCREMENT,
    Username  TEXT     NOT NULL UNIQUE,
    Email     TEXT,
    Password  TEXT     NOT NULL,
    FirstName TEXT,
    LastName  TEXT,
    Role      INTEGER  NOT NULL,
    CreatedAt DATETIME NOT NULL,
    UpdatedAt DATETIME NOT NULL
);

-- Username: admin
-- Password: admin
INSERT INTO Users (Username, Email, Password, FirstName, LastName, Role, CreatedAt, UpdatedAt)
VALUES ('admin', null, '$2a$10$Z13RQlu6HdKSW41rJsz7Ju5NZ0VMyUdm6YZMr0wjJqW955qd2pzx2',
        null, null, 1, '1731426770000', '1731426770000');

CREATE TABLE Products
(
    ID                  INTEGER PRIMARY KEY AUTOINCREMENT,
    Name                TEXT     NOT NULL UNIQUE,
    Slug                TEXT     NOT NULL UNIQUE,
    Description         TEXT,
    Logo                TEXT,
    Private             BOOLEAN,
    VersionFormat       TEXT,
    SetupGuideCompleted BOOLEAN DEFAULT FALSE,
    CreatedAt           DATETIME NOT NULL,
    UpdatedAt           DATETIME NOT NULL
);

CREATE TABLE Channels
(
    ID        INTEGER PRIMARY KEY AUTOINCREMENT,
    Name      TEXT NOT NULL,
    ProductID INTEGER
        CONSTRAINT Channels_Products_ID_fk REFERENCES Products,
    Closed    BOOLEAN DEFAULT FALSE
);

CREATE UNIQUE INDEX Channels_ProductID_Name_uindex
    ON Channels (ProductID, Name);

CREATE TABLE Apps
(
    ID        INTEGER PRIMARY KEY AUTOINCREMENT,
    ProductID INTEGER
        CONSTRAINT Apps_Products_ID_fk REFERENCES Products,
    Name      TEXT     NOT NULL,
    OS  TEXT,
    CreatedAt DATETIME NOT NULL,
    UpdatedAt DATETIME NOT NULL
);

CREATE UNIQUE INDEX Apps_ProductID_OS_uindex
    ON Apps (ProductID, OS);

CREATE TABLE Builds
(
    ID        INTEGER PRIMARY KEY AUTOINCREMENT,
    AppID     INTEGER
        CONSTRAINT Builds_App_ID_fk REFERENCES Apps,
    Number    TEXT     NOT NULL,
    Version   TEXT     NOT NULL,
    CreatedAt DATETIME NOT NULL,
    UpdatedAt DATETIME NOT NULL
);

CREATE UNIQUE INDEX Builds_AppID_Number_uindex
    ON Builds (AppID, Number);

CREATE TABLE BuildMetadata
(
    BuildID INTEGER
        CONSTRAINT BuildMetadata_Build_ID_fk REFERENCES Builds,
    Key     TEXT NOT NULL,
    Value   TEXT NOT NULL,
    CONSTRAINT BuildMetadata_pk PRIMARY KEY (BuildID, Key)
);

CREATE TABLE Artifacts
(
    ID           INTEGER PRIMARY KEY AUTOINCREMENT,
    BuildID      INTEGER
        CONSTRAINT Artifacts_Build_ID_fk REFERENCES Builds,
    Md5          TEXT    NOT NULL,
    Sha256       TEXT    NOT NULL,
    Sha512       TEXT    NOT NULL,
    Filename     TEXT    NOT NULL,
    MimeType     TEXT    NOT NULL,
    Size         INTEGER NOT NULL,
    Path         TEXT    NOT NULL,
    Architecture TEXT    NOT NULL
);

CREATE UNIQUE INDEX Artifacts_BuildID_Architecture_uindex
    ON Artifacts (BuildID, Architecture);
