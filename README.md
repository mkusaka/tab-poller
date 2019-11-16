# tab-poller
polling running chrome tabs.

# useage
1. clone this repo
1. run server: `make server`
1. get list:  `http://localhost:30303`
1. activate tab: `http://localhost:30303?url=https://github.com/mkusaka/tab-poller` // url must be url encoded

# feature
- [x] polling running chrome tabs & provide it via http api.
- [x] add activate tab http api (parameter: tab's (url encoded) url).
- [ ] manage tab data in-memory.
- [ ] run via cli & run background.
- [ ] make settings configurable.
    - polling duration
    - port
- [ ] stop via cli or http request.
- [ ] terminate this application via http request.
