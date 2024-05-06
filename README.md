# Data-Pipeline
Data Pipeline Project in Go

Started with code-heim's data pipeline (https://github.com/code-heim/go_21_goroutines_pipeline), added benchmarking, some unit testing, along with error checking for file input/output validation. Files that are not found are removed and not processed instead of causing the entire program to reach panic state. Later iterations should address panics and establish a test behavior to process files which can be processed versus causing the entire program to crash. Replaced the four images given in the intial repo to some of my own choosing. Also added additional folder for unit tests and benchmarking.

Benchmarking results in program finishing in roughly 0.538s for processing four images.
