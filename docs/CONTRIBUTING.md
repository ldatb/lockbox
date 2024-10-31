# Contributing to Lockbox

Thank you for considering contributing to **Lockbox**! We appreciate your help in making this project even better. This guide will walk you through our contribution process and provide you with best practices to help your contributions be successful.

## Code of Conduct

By participating in this project, you agree to abide by our [Code of Conduct](CODE_OF_CONDUCT.md). Please ensure all interactions are respectful and constructive.

## How to Contribute

There are several ways to contribute to Lockbox:

1. **Report Bugs**: Use [GitHub Issues](https://github.com/ldatb/lockbox/issues) to report bugs, providing as much detail as possible.
2. **Propose Features**: Open a GitHub Issue to propose new features and discuss potential ideas with maintainers.
3. **Improve Documentation**: Help make Lockbox’s documentation clearer and more comprehensive.
4. **Submit Code**: Add new features, fix bugs, or improve existing functionality.

## Setting Up Your Environment

To get started with development, you’ll need to set up a local environment:

1. **Fork the repository** and clone it locally:

```bash
git clone https://github.com/your-username/lockbox.git
cd lockbox
```

2. **Install dependencies**:

```bash
make mod
```

3. **Run tests** to ensure everything is working:

```bash
make test
```

Check out [HACKING.md](HACKING.md) for more detailed setup and configuration instructions.

## Making Changes

1. **Create a new branch** for each feature, bug fix, or documentation update:

```bash
git checkout -b feature/your-feature-name
```

2. **Write clear and concise code** that aligns with project conventions. Use comments to clarify complex sections where needed.

3. **Follow best practices** in Go and adhere to code style guidelines. Use `go fmt` to format your code consistently.

## Writing Tests

All new features and bug fixes must include corresponding tests:

- Place **unit tests** in the relevant package’s directory (e.g., `pkg/crypto/`).
- Place **integration tests** in the `test/integration/` directory.
- Run the tests using:

```bash
go test -v ./...
```

Testing is essential to ensure Lockbox’s security and reliability.

## Submitting a Pull Request

Once you’re ready to submit your changes:

1. **Push your branch** to your fork:

```bash
git push origin feature/your-feature-name
```

2. **Open a Pull Request (PR)** from your branch on GitHub, targeting the `main` branch of the Lockbox repository.

3. **Describe your changes** clearly in the PR description. Include:
   - A brief summary of what you did
   - Any relevant issue numbers (e.g., "Fixes #123")
   - Screenshots or logs (if applicable)

4. **Respond to review feedback**: A maintainer will review your PR. Please address feedback promptly and discuss any changes or questions as needed.
