# Golang Tutorial

- [x] Create a binary for intruder-alert
- [ ] Create Golang project structure (cmd, pkg, etc)
- [ ] Configure a connection to a pg2pubsub subscription
  - [ ] Create the subscription in terraform in advance
- [ ] Confirm we can see changes as they come through the subscription
- [ ] Create an alert that logs whenever an org creates more than 5 payments/min

After this is complete, we can work on refactoring to enable unit tests and
create a sensible pluggable structure for general alerts.
