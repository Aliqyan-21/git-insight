## Git Insight

git-insight is a lightweight and easy-to-use command-line tool designed to fetch and display recent GitHub events associated with a given `username`. It provides detailed insights into key GitHub activities, such as push events, pull requests, issue comments, repository creations, and more. Whether you're tracking your own activities or monitoring the actions of other developers, git-insight helps you gain quick access to meaningful interactions on GitHub, in the terminal itself!.

## Preview

https://github.com/user-attachments/assets/3b31c5f3-270d-46ad-afa5-9965c7983ecd

## Installation

To install git-insight, follow these steps:

### Prerequisites

- Go 1.18 or later is required to run or build this application.

### Building from Source

1. Clone the repository:

```bash
git clone git@github.com:Aliqyan-21/git-insight.git
cd git-insight
```

2. Build the application:

```bash
go build cmd/main.go && mv main git-insight
```

> Run the application

```bash
./git-insight <username>
```

Alternatively, you can install it globally with:

```bash
go install
```

> Run the application

```bash
git-insight <username>
```

License
git-insight is licensed under the MIT License.
