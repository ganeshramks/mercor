<h3>How to centralise this abstraction so that language change does not increase the work for engineers?</h3>

1. We could use an exclusive scd microservice. This will be a standalone service which exposes API endpoints for the operations such as read/update job, timelog etc... 
2. A significant benefit is that a client written in any language can interact with the service. 
3. All the scd logic is maintained in one place.
4. The DB on this service can be anything 
