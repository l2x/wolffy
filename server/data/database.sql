create table `cluster` 
    -- --------------------------------------------------
    --  Table Structure for `github.com/l2x/wolffy/server/models.Cluster`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `cluster` (
        `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
        `name` varchar(255) NOT NULL,
        `tags` varchar(255) NOT NULL,
        `machine` varchar(255) NOT NULL,
        `note` varchar(255) NOT NULL,
        `created` datetime NOT NULL,
        `modified` datetime NOT NULL,
        UNIQUE (`name`)
    );

create table `deploy` 
    -- --------------------------------------------------
    --  Table Structure for `github.com/l2x/wolffy/server/models.Deploy`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `deploy` (
        `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
        `pid` integer NOT NULL,
        `commit` varchar(255) NOT NULL,
        `status` integer NOT NULL,
        `created` datetime NOT NULL,
        `modified` datetime NOT NULL
    );
    CREATE INDEX `deploy_pid` ON `deploy` (`pid`);

create table `deploy_history` 
    -- --------------------------------------------------
    --  Table Structure for `github.com/l2x/wolffy/server/models.DeployHistory`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `deploy_history` (
        `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
        `did` integer NOT NULL,
        `status` integer NOT NULL,
        `note` varchar(255) NOT NULL,
        `created` datetime NOT NULL,
        `modified` datetime NOT NULL
    );
    CREATE INDEX `deploy_history_did` ON `deploy_history` (`did`);

create table `machine` 
    -- --------------------------------------------------
    --  Table Structure for `github.com/l2x/wolffy/server/models.Machine`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `machine` (
        `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
        `ip` varchar(255) NOT NULL,
        `port` varchar(255) NOT NULL,
        `note` varchar(255) NOT NULL,
        `token` varchar(255) NOT NULL,
        `status` integer NOT NULL,
        `created` datetime NOT NULL,
        `modified` datetime NOT NULL,
        UNIQUE (`ip`)
    );

create table `project` 
    -- --------------------------------------------------
    --  Table Structure for `github.com/l2x/wolffy/server/models.Project`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `project` (
        `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
        `name` varchar(255) NOT NULL,
        `path` varchar(255) NOT NULL,
        `push_path` varchar(255) NOT NULL,
        `tags` varchar(255) NOT NULL,
        `note` varchar(255) NOT NULL,
        `created` datetime NOT NULL,
        `modified` datetime NOT NULL,
        UNIQUE (`name`)
    );

create table `project_cluster` 
    -- --------------------------------------------------
    --  Table Structure for `github.com/l2x/wolffy/server/models.ProjectCluster`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `project_cluster` (
        `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
        `pid` integer NOT NULL,
        `cid` integer NOT NULL,
        `custom_machine` varchar(255) NOT NULL,
        `bshell` varchar(255) NOT NULL,
        `eshell` varchar(255) NOT NULL,
        `note` varchar(255) NOT NULL,
        `created` datetime NOT NULL,
        `modified` datetime NOT NULL
    );
    CREATE INDEX `project_cluster_pid` ON `project_cluster` (`pid`);

