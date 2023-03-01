describe('The Login Page', () => {
  it('Opens login page of local host', () => {
    cy.visit('http://localhost:4200')
  })
  
  it('Opens login page of local host, then clicks sign up button', () => {
    cy.visit('http://localhost:4200')

    cy.contains('Sign up').click()
  })

  it('Fills out username and password, then clicks login', () => {
    cy.visit('http://localhost:4200')

    cy.get("#Username").type("TESTUSER")

    cy.get("#Password").type("TESTPASS")

    cy.get('#submit').click()
  })
})