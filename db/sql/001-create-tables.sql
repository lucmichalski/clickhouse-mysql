---- drop ----
DROP TABLE IF EXISTS `test_table`;

---- create ----
CREATE TABLE `logs` (
  `id1` int(10) unsigned NOT NULL,
  `id2` int(10) unsigned NOT NULL,
  `id3` int(10) unsigned NOT NULL,
  `id4` int(10) unsigned NOT NULL,
  `id5` int(10) unsigned NOT NULL,
  `id6` int(10) unsigned NOT NULL,
  `id7` int(10) unsigned NOT NULL,
  `id8` int(10) unsigned NOT NULL,
  `id9` int(10) unsigned NOT NULL,
  `log_time` datetime NOT NULL,
  `request` int(10) unsigned DEFAULT NULL,
  `imp` int(10) unsigned DEFAULT NULL,
  `click` int(10) unsigned DEFAULT NULL,
  `revenue` decimal(20,6) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id1`,`id2`,`id3`,`id4`,`id5`,`id6`,`id7`,`id8`,`id9`,`log_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
