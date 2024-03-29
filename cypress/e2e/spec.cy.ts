describe('template spec', () => {
  it('Local Host Opened', () => {
    cy.visit('http://localhost:4200/');
    cy.url().should('eq', 'http://localhost:4200/home')

    //Cypress.on('uncaught:exception', (err, runnable) => { return false; });
  })

  it('Test Home Routing', () => {
    cy.visit('http://localhost:4200/');
    cy.get('#home-button').click();
    cy.url().should('eq', 'http://localhost:4200/home')
  })
  it('Test Login Routing', () => {
    cy.visit('http://localhost:4200/');
    cy.get('#login-button').click();
    cy.url().should('eq', 'http://localhost:4200/login')
  })
  it('Test Sign Up Routing', () => {
    cy.visit('http://localhost:4200/');
    cy.get('#signup-button').click();
    cy.url().should('eq', 'http://localhost:4200/signup')
  })

  it('Test Sign Up Functionality', () => {
    cy.visit('http://localhost:4200/');
    cy.get('#signup-button').click();
    cy.get('#Username').type('test');
    cy.get('#Password').type('testpass');
    cy.get('#submit').click();
  })

  it('Test Login Functionality', () => {
    cy.visit('http://localhost:4200/');
    cy.get('#login-button').click();
    cy.get('#Username').type('test');
    cy.get('#Password').type('test');
    cy.get('#submit').click();
    cy.url().should('eq', 'http://localhost:4200/review')
  })

  it('Test Sign Up then Login', () => {
    cy.visit('http://localhost:4200/');
    cy.get('#signup-button').click();
    cy.get('#Username').type('test');
    cy.get('#Password').type('test');
    cy.get('#submit').click();

    cy.get('#login-button').click();
    cy.get('#Username').type('test');
    cy.get('#Password').type('test');
    cy.get('#submit').click();
    cy.url().should('eq', 'http://localhost:4200/review');
  })


  it('Test Login then Review', () => {
    cy.visit('http://localhost:4200/');
    cy.get('#login-button').click();
    cy.get('#Username').type('test');
    cy.get('#Password').type('test');
    cy.get('#submit').click();

    cy.get('#search-input').type('Taylor Swift');
    cy.get('#Confirm-button').click();
    cy.get('#review-input').type('Incredible Music!');
    cy.get('.star-rating > :nth-child(5)').click();
    cy.get('#submit-button').click();
  })

  it('Review Functionality', () => {
    cy.visit('http://localhost:4200/');
    cy.get('#login-button').click();
    cy.get('#Username').type('test');
    cy.get('#Password').type('test');
    cy.get('#submit').click();

    cy.get('#search-input').type('Taylor Swift');
    cy.get('#Confirm-button').click();
    cy.get('#review-input').type('Cool Beans!');
    cy.get('.star-rating > :nth-child(5)').click();
    cy.get('#submit-button').click();
  })

})
