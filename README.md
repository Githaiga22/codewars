# Code Wars

Welcome to the Code Wars repository! This is a collection of various coding challenges and tasks that I am solving as part of my practice. The repository contains solutions in Go (Golang), with a focus on writing clean, efficient, and well-documented code.
Repository Structure

The repository is organized as follows:
```bash
codewars/
│
├── README.md           # Overview of the repository
├── <challenge-name>/   # Directory for each challenge or task
│   ├── main.go         # Main solution file for the challenge
│   ├── *_test.go       # Test cases for the challenge (if applicable)
│   └── README.md       # Details specific to the challenge (description, approach, etc.)
│
└── others/             # Additional scripts or utilities used in challenges
```
Challenge Directories

Each challenge is contained in its own directory. For example:
```bash
codewars/
├── split-strings/
│   ├── main.go         # Solution for the "Split Strings" challenge
│   ├── main_test.go    # Test cases for the solution
│   └── README.md       # Description and explanation of the challenge
```
### Main Files

- main.go: Contains the primary solution to the challenge.
- *_test.go: Contains test cases to validate the solution.
-  README.md: Explains the challenge, the approach taken to solve it, and any edge cases or optimizations considered.

## Running the Solutions

To run a solution, navigate to the respective challenge directory and execute the following command:
```bash
go run main.go
```
To run the test cases:
```bash
go test
```
## Contributing

Contributions are welcome! If you have suggestions, improvements, or new challenges you'd like to add, feel free to submit a pull request.

- Fork the repository.
-   Create a new branch (git checkout -b feature/your-feature).
-   Make your changes.
-   Commit your changes (git commit -am 'Add your feature').
-   Push to the branch (git push origin feature/your-feature).
-   Create a new Pull Request.

## License

This repository is licensed under the MIT License. See the LICENSE file for more details.