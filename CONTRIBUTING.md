# Contribution Guide and Git Workflow

This document explains the Git workflow, code contribution guidelines, and release process to ensure a smooth and efficient experience for everyone.

## Branching Strategy

We use a structured Git branching strategy to ensure that development remains smooth and efficient. Our key branches are as follows:

### Main Branch (`main`)
- **Purpose**: The `main` branch contains **production-ready** code.
- **Usage**: Only stable and thoroughly tested code is merged into `main`.
- **Releases**: Each release is tagged in this branch (e.g., `v1.0`, `v1.1`).

### Develop Branch (`develop`)
- **Purpose**: This is the **primary branch** for ongoing development.
- **Usage**: All feature branches are merged into `develop`. This branch contains the latest work and is subject to frequent updates.

### Feature Branches (`feature/*`)
- **Purpose**: Used for **developing new features** or **fixing bugs**.
- **Usage**: Feature branches are created from `develop` and should be merged back into `develop` after completion and code review.
- **Naming convention**: `feature/<feature-name>`.

### Bug Branches (bugfix/*)
- **Purpose**: Used for fixing bugs.
- **Usage**: Bugfix branches are created from `develop` (or `main` in case of critical bugs) and merged back into `develop` after the fix is complete and reviewed.
- **Naming convention**: `bugfix/<bug-description>.`

### Release Branches (`release/*`)
- **Purpose**: **Stabilizing a new release** and preparing deployment artifacts.
- **Usage**: A release branch is created from `develop` when the software is feature-complete and stable. This branch is used to fix bugs, polish features, and finalize documentation.
- **Naming convention**: `release/<version-number>`.

### Hotfix Branches (`hotfix/*`)
- **Purpose**: For **fixing critical bugs** or security issues in a production release.
- **Usage**: Hotfix branches are created from `main` and merged back into both `main` and `develop` to ensure that the fix is included in future releases.
- **Naming convention**: `hotfix/<hotfix-name>`.

## Workflow Process

1. **Feature Development**
   - Develop features or bug fixes in feature / bug branches based on `develop`.
   - Once complete, feature / bug branches are merged back into `develop`.

2. **Release Process**
   - A **release branch** is created from `develop` when the codebase is stable and ready for a release.
   - Final touches, bug fixes, and deployment preparations happen on the release branch.
   - The release branch is merged into `main` and tagged with a version number.

3. **Hotfixes**
   - Critical fixes are handled through **hotfix branches** off `main`.
   - Hotfixes are merged into both `main` and `develop` to ensure they are included in future releases.

## Versioning and Tagging

We use **semantic versioning** to track releases:
- **Major**: Breaking changes (e.g., `v2.0`).
- **Minor**: New features, backward-compatible (e.g., `v1.1`).
- **Patch**: Bug fixes and security patches (e.g., `v1.0.1`).

Tagging a release:

```bash
git tag -a v<version-number> -m "Version <version-number>"
git push origin main --tags
```

## Deployment Artifacts

Each release includes the following artifacts for self-hosting:

- **Helm Charts**: For easy deployment on Kubernetes.
- **Docker Images**: Pre-built Docker images for running the Go application.
- **Documentation**: Detailed setup instructions for deploying the self-hosted version.

These artifacts are created during the **release phase** and made available with each version tag.

## Submitting Pull Requests

1. **Fork the Repository**: Start by forking the repository and cloning it to your local machine.

```bash
git clone git@gitlab.com:xrs-cloud/lockbox/core.git
```

2. **Create a Feature Branch**: Develop your changes in a feature branch.

```bash
git checkout -b feature/<feature-name>
```

3. **Commit Your Changes**: Write meaningful commit messages and ensure the commit history is clean.

```bash
git commit -m "Add feature <description>"
```

4. **Push Your Branch**: Push your feature branch to your forked repository.

```bash
git push origin feature/<feature-name>
```

5. **Create a Pull Request**: Open a pull request from your feature branch to `develop`. Provide a detailed description of the changes, along with links to relevant issues.

### Code Review Process

- All pull requests are subject to code review.
- At least one other contributor must approve your PR before it is merged.
- Make sure your changes are tested locally and provide unit tests where applicable.

## Code Style Guidelines

To ensure code consistency, please follow these guidelines:

1. **Use Clear and Descriptive Names**: Use meaningful variable, function, and class names.
2. **Write Tests**: For each new feature or bug fix, ensure there are adequate unit or integration tests.
3. **Follow Linting Rules**: Use the provided linting configurations to keep code style consistent across the codebase.
