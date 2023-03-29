describe('Verify Browser Stack Home Page', () => {

      it('Successfully Load the Local Host', () => {
        cy.visit('http://localhost:4200/');
        Cypress.on('uncaught:exception', (err, runnable) => { return false; });
        /*cy.get('#uname').should('be.visible');*/

      })

  })
