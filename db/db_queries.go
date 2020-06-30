package db

const sqlStatementGetMeasurements = `
SELECT object, received_at
FROM device_up
WHERE f_port = 1
AND object @> '{"func": 1}'
AND dev_eui=$1
AND ($2::TIMESTAMP IS NULL OR received_at >= $2)
AND ($3::TIMESTAMP IS NULL OR received_at <= $3)
ORDER BY received_at DESC
`

const sqlStatementDeleteMeasurements = `
DELETE FROM device_up
WHERE f_port = 1
AND object @> '{"func": 1}'
AND dev_eui=$1
AND ($2::TIMESTAMP IS NULL OR received_at >= $2)
AND ($3::TIMESTAMP IS NULL OR received_at <= $3)
`

const sqlStatementMeasurementOverviewDetails = `
SELECT DISTINCT d_max.max_received_at, d_min.min_received_at, d_num_meas.number_of_measurements
FROM device_up d
INNER JOIN
(
	SELECT dev_eui, MAX(received_at) as max_received_at
	FROM device_up
	GROUP BY dev_eui
) d_max on d_max.dev_eui = d.dev_eui
INNER JOIN
(
	SELECT dev_eui, MIN(received_at) as min_received_at
	FROM device_up
	GROUP BY dev_eui
) d_min on d_min.dev_eui = d.dev_eui
INNER JOIN
(
	SELECT dev_eui, COUNT(*) as number_of_measurements
	FROM device_up
	WHERE f_port = 1
	AND object @> '{"func": 1}'
	GROUP by dev_eui
) d_num_meas ON d_num_meas.dev_eui = d.dev_eui
WHERE d.dev_eui=$1
`

const sqlStatementNewestDeviceInfo = `
SELECT object
FROM device_up
WHERE dev_eui=$1
AND object @> '{"func": 12}'
ORDER BY received_at DESC
LIMIT 1
`

const sqlStatementTableExists = `SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_name = $1);`
