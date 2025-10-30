# Gola Language

**Gola** is a cute and simple programming language compiler designed to be easy to use. This guide will help you install and get started with Gola on your system.
![Gola Language](https://github.com/felixoder/gola-language/blob/main/assets/gola.png)

## Using Gola
![Gola Language](https://github.com/felixoder/gola-language/blob/main/assets/demo.png)


## Features
- Lightweight and easy to use
- Supports Linux, macOS
- Built for simplicity and speed

## Installation

You can install Gola on your system by following the steps below:
###### The installation is for the RECENT version of Gola, if you want to install the OLD version, please click here

### Step 1: Download the Gola Compiler

#### For Linux (x86_64)
```bash
curl -L -o gola-linux-x86_64.tar.gz https://github.com/felixoder/gola-language/releases/download/v2.0.0/gola-linux-x86_64.tar.gz


```
#### For macOS (x86_64)

```bash
curl -L -o gola-darwin-x86_64.tar.gz https://github.com/felixoder/gola-language/releases/download/v2.0.0/gola-darwin-x86_64.tar.gz
```
#### Step 2: Extract the Files

```bash
tar -xvzf gola-linux-x86_64.tar.gz
```
##### or for macOS
```bash
tar -xvzf gola-darwin-x86_64.tar.gz
```

#### Step 3: Step 3: Install Gola

```bash
sudo mv ./v2.0.1/gola-linux-x86_64 /usr/local/bin/gola
```
```bash
sudo chmod +x /usr/local/bin/gola
```

#### Step 4: Verify the Installation

```bash
gola -v
```
