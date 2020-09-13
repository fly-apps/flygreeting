## Continuous Deployment to Fly with Codeship Pro

- What is CD? Why it’s important? (Just briefly)
- What is Codeship? Why use it for CI/CD?
- Why host on the edge? Why Fly? (Again, briefly for context)

## How to Deploy Using Codeship Pro
- Fork/clone flygreeting
- Set up Codeship on the new repo
  - Get AES key, save as codeship.aes
  - Update .gitignore file
- New Dockerfile for running tests
- Deploy.sh script
- Add codeship-services.yml, codeship-steps.yml
- Remove the toml file, flyctl init
- Save FLYCTL_AUTH_TOKEN to env file
- Using jet to encrypt the env file
- Commit new files, test it out

## Conclusion

