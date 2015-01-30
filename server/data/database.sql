create table `project` 
    -- --------------------------------------------------
    --  Table Structure for `github.com/l2x/wolffy/server/models.Project`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `project` (
        `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
        `pid` integer NOT NULL,
        `name` varchar(255) NOT NULL,
        `path` varchar(255) NOT NULL,
        `note` varchar(255) NOT NULL,
        `created` datetime NOT NULL
    );

create table `cluster` 
    -- --------------------------------------------------
    --  Table Structure for `github.com/l2x/wolffy/server/models.Cluster`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `cluster` (
        `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
        `name` varchar(255) NOT NULL,
        `room` varchar(255) NOT NULL,
        `machine` varchar(255) NOT NULL,
        `note` varchar(255) NOT NULL,
        `created` datetime NOT NULL
    );

create table `deploy` 
    -- --------------------------------------------------
    --  Table Structure for `github.com/l2x/wolffy/server/models.Deploy`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `deploy` (
        `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
        `pid` integer NOT NULL,
        `commit` varchar(255) NOT NULL,
        `created` datetime NOT NULL
    );

create table `product` 
    -- --------------------------------------------------
    --  Table Structure for `github.com/l2x/wolffy/server/models.Product`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `product` (
        `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
        `name` varchar(255) NOT NULL,
        `note` varchar(255) NOT NULL,
        `created` datetime NOT NULL
    );

