describe('Initial Visit', () => {
  it('Visits the initial project page', () => {
    cy.viewport(1024, 768)
    cy.visit('/')
    cy.contains('Insert Company Name')
  })
})

describe('Click for Login and attempts to log in', () => {
  it('Clicks the login button', () => {
    //Without this there are some button visibility issues?
    cy.viewport(1024, 768)
    cy.visit('/')
    cy.contains('Login').click()
    cy.url().should('include', '/login')
    //Type credentials
    cy.get('input[name="username"]')
      .type('testname@fake.com')
      .should('have.value','testname@fake.com')
    cy.get('input[name="password"]')
      .type('password1')
      .should('have.value','password1')
  })
})

describe('Views the tables', () => {
  it('Visits the table tabs', () => {
    cy.viewport(1024, 768)
    cy.visit('/')
    cy.contains('Test User Page').click()
    cy.url().should('include', '/browse-contracts')
    //Had to do get with selector for this to work?
    //Look at all three tables (messages then notifications then contracts)
    cy.get('body > app-root > app-browse-contracts > div > mat-sidenav-container > mat-sidenav-content > mat-nav-list > a:nth-child(2) > span').click()
    cy.contains('Recipients')
    cy.get('body > app-root > app-browse-contracts > div > mat-sidenav-container > mat-sidenav-content > mat-nav-list > a:nth-child(1) > span').click()
    cy.contains('Required Action')
    cy.get('body > app-root > app-browse-contracts > div > mat-sidenav-container > mat-sidenav-content > mat-nav-list > a:nth-child(3) > span').click()
    cy.contains('Involved parties')
  })
})
