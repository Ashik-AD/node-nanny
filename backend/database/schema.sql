CREATE TABLE IF NOT EXISTS projects (
        id TEXT PRIMARY KEY,
        name TEXT NOT NULL UNIQUE,
        path TEXT NOT NULL UNIQUE,
        detail TEXT NOT NULL,
        files TEXT,
        dependencies TEXT,
        subprojects TEXT,
        logs TEXT,

        FOREIGN KEY(detail) REFERENCES details(id),
        FOREIGN KEY(files) REFERENCES files(pid),
        FOREIGN KEY(dependencies) REFERENCES dependencies(pid),
        FOREIGN KEY(subprojects) REFERENCES subprojects(root_id),
        FOREIGN KEY(logs) REFERENCES logs(pid)
);

CREATE TABLE IF NOT EXISTS subprojects (
        id TEXT PRIMARY KEY,
        root_id TEXT NOT NULL,
        name TEXT NOT NULL,
        path text NOT NULL,
        detail TEXT NOT NULL,
        files TEXT,
        dependencies TEXT,
        logs TEXT,

        FOREIGN KEY(detail) REFERENCES details(id),
        FOREIGN KEY(files) REFERENCES files(pid),
        FOREIGN KEY(dependencies) REFERENCES dependencies(pid),
        FOREIGN KEY(logs) REFERENCES logs(pid)
);

CREATE TABLE IF NOT EXISTS details (
        id TEXT PRIMARY KEY,
        pid TEXT NOT NULL,
        total_modules INTEGER,
        storage INTEGER,
        dependencies INTEGER,
        dev_depedencies INTEGER,
        outdated INTEGER,
        warning INTEGER,
        dangered INTEGER
);

CREATE TABLE IF NOT EXISTS dependencies (
    pid TEXT NOT NULL,
    pkg_id TEXT NOT NULL,
    installed_ver TEXT NOT NULL,

    PRIMARY KEY(pid, pkg_id),
    FOREIGN KEY(pkg_id) REFERENCES pkgdetails(id)
);

CREATE TABLE IF NOT EXISTS pkgdetails (
        id TEXT PRIMARY KEY,
        name TEXT NOT NULL UNIQUE,
        version TEXT,
        npm_link TEXT,
        home_link TEXT,
        repo_link TEXT,
        size REAL
);

CREATE TABLE IF NOT EXISTS files (
        id TEXT PRIMARY KEY,
        pid TEXT NOT NULL,
        name TEXT NOT NULL,
        path TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS logs (
        id TEXT PRIMARY KEY,
        pid TEXT NOT NULL,
        text TEXT NOT NULL,
        messages TEXT,
        status TEXT,
        timestamp DATETIME,

        FOREIGN KEY(messages) REFERENCES log_messages(id)
);

CREATE TABLE IF NOT EXISTS log_messages (
        id TEXT PRIMARY KEY,
        text TEXT NOT NULL,
        status TEXT,
        timestamp DATETIME
);
