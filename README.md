# DDD & Event sourcing workshop

### What we will try to implement
* an event-sourced aggregate: Person
* some value objects for a Person: Name, EmailAddress, Address
* some domain events: PersonRegistered, PersonEmailAddressConfirmed, PersonAddressAdded, PersonAddressChanged
* the methods in the Person aggregate which will cause the above events
* we'll try to work test-driven as much as possible

### Let's apply 
 * DDD principles and 
 * make the model Event sourced

### My choice of programming language
 * Golang

# Steps
 * Create `Name` value object and tests (I will not mention tests in the following steps, just create them always ;)
 * Create simple `Email` value object, only with the "email" value (confirm will come later)
 * Create `Address` value object with basic fields (CountryCode, PostalCode,City,Street,HouseNumber)
 * Create `Person` entity (aggregate root) with basic method `Register`
 * Create domain event `PersonRegistered`
 * Follow with methods `RecordedEvents`, `recordThat` and `apply`
 * Add new behaviour to the `Email` - "confirm email" 
   * adapt value object and also the aggregate
   * create domain event `EmailConfirmed`
 * 
 
### TODOs
 * Finish `Address` events 
   * AddressAdded
   * AddressChanged
  
  
# Notes / findings
 * Visibilities in Golang are not that powerful as in other languages
 