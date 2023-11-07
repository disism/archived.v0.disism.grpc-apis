# Archived!

Dependency management with `buf.build`

## VERSION



## GIT

Triggering make version on commit before commit is used to commit the current development version of the buf

```
#!/bin/sh

# Run buf lint
echo "Running buf lint..."
buf lint
if [ $? -ne 0 ]; then
  echo "Error: buf lint failed. Aborting script."
  exit 1
fi
echo "buf lint successful."

# Run buf format
echo "Running buf format..."
buf format
if [ $? -ne 0 ]; then
  echo "Error: buf format failed. Aborting script."
  exit 1
fi
echo "buf format successful."

# Run make version
echo "Running make version..."
make version
if [ $? -ne 0 ]; then
  echo "Error: make version failed. Aborting script."
  exit 1
fi
echo "make version successful."

# Stage all changes
git add .

```
