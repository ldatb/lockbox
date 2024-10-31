# HACKING Lockbox

Welcome to the **Lockbox** hacking guide! This document provides instructions for developers looking to contribute to Lockbox, understand its codebase, and get started with making improvements or changes.

## Getting Started

### Prerequisites

- **Go 1.23** or higher
- Docker (if you plan to test in a containerized environment)
- Familiarity with Git and GitHub

### Initial Setup

1. **Clone the repository**:

```bash
git clone https://github.com/ldatb/lockbox.git
cd lockbox
```

2. **Install dependencies**:

```bash
make mod
```

3. **Build the project**:

```bash
make dev
```

4. **Run tests** to verify the setup:

```bash
make test
```

## Development Workflow

1. **Create a new branch** for each feature or bug fix:

```bash
git checkout -b feature/new-awesome-feature
```

2. **Make changes** in your branch. Always keep the changes modular to facilitate reviews.

3. **Test your changes** locally.

4. **Commit your changes** with descriptive messages:

```bash
git commit -m "Add feature to integrate XYZ functionality"
```

5. **Push your branch** to GitHub:

```bash
git push origin feature/new-awesome-feature
```

6. **Submit a Pull Request (PR)** to the main branch.

## Best Practices

- **Write Clear and Modular Code**: Keep functions short and focused on a single task.
- **Error Handling**: Handle errors explicitly, and provide meaningful error messages.
- **Code Consistency**: Follow the project's Go style and linting rules
- **Security**: Avoid hardcoding credentials or secrets, and use environment variables wherever possible.

## Contributing

Lockbox is open to contributions! To start:

1. Check for existing issues or create a new one if your change is significant.
2. Discuss major changes with maintainers before beginning work.
3. Write tests and documentation for new features.
4. Follow the [contributing guidelines](CONTRIBUTING.md) for details.

## Getting Help

If you run into issues, need clarification, or have ideas, reach out via:

- GitHub Issues: [Create an Issue](https://github.com/ldatb/lockbox/issues)

Happy Hacking! Thank you for your interest in contributing to **Lockbox**.
