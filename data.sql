INSERT INTO tests(status, title, details) VALUES
('Passed','Database connection check', '{"Verify that the application can successfully connect to the database using the configured credentials.",
"Ensure that connection pooling works correctly and no errors occur during multiple concurrent connections.",
"Test the database timeout behavior when the database is slow or unreachable."}'),

('Failed','Login with invalid credentials', '{"Attempt to login using incorrect username and/or password combinations.",
"Check that the system displays appropriate error messages without revealing sensitive information.",
"Ensure that account lockout policies are enforced after multiple failed attempts."}'),

('Open','Search performance test', '{"Measure the response time of the search functionality under normal load conditions.",
"Validate that search results are accurate and relevant according to the query parameters.",
"Test search performance under high load and ensure system stability and acceptable response times."}'),

('Blocked','Stripe payment validation', '{"Verify that the payment process integrates correctly with Stripe''s API in test mode.",
"Check that successful transactions update order status correctly in the application.",
"Handle error scenarios such as declined payments or network failures and ensure proper user feedback."}'),

('Open','CSV data export', '{"Ensure that users can export data to CSV format from different sections of the application.",
"Verify that the CSV files contain correct headers and all expected data fields.",
"Test large data exports for performance and file integrity, ensuring no data is truncated or corrupted."}');
