

describe('Initial Visit', () => {
  it('Visits the initial project page', () => {
    cy.viewport(1024, 768)
    cy.visit('/')
    cy.contains('Insert Company Name')
  })
})

describe('Creating an account works', () => {
  it('Make a fake account', () => {
    //Without this there are some button visibility issues?
    cy.viewport(1024, 768)
    cy.visit('/')
    cy.contains('Login').click()
    cy.url().should('include', '/login')
    cy.contains('Create Account').click()
    //Type credentials
    cy.get('input[name="name"]')
      .type('Fake Name')
      .should('have.value','Fake Name')
    cy.get('input[name="username"]')
      .type('fakeuser123')
      .should('have.value','fakeuser123')
    cy.get('input[name="email"]')
      .type('fakeemail123@notreal.com')
      .should('have.value','fakeemail123@notreal.com')
    cy.get('input[name="password"]')
      .type('password1')
      .should('have.value','password1')
    cy.contains('Create Account').click();
    cy.scrollTo('top');
      //Click alert?
    cy.contains('User Creation Successful');
    cy.wait(1000);
  })
})

describe('Log in as fake person', () => {
  it('Log in as fake person', () => {
    //Without this there are some button visibility issues?
    cy.viewport(1024, 768)
    cy.visit('/')
    cy.contains('Login').click()
    cy.url().should('include', '/login')
    //Type credentials
    cy.get('input[name="email"]')
      .type('fakeemail123@notreal.com')
      .should('have.value','fakeemail123@notreal.com')
    cy.get('input[name="password"]')
      .type('password1')
      .should('have.value','password1')
    cy.get('body > app-root > app-login-form > section > div > div > div.columns.is-centered > div > form > div.field.is-grouped > div:nth-child(1) > button').click();
      //Click alert?
    cy.scrollTo('top');
    cy.contains('User Login Successful');
    cy.wait(1000);
  })
})

describe('Views the tables', () => {
  it('Visits the table tabs', () => {
    cy.viewport(1024, 768)
    cy.visit('/')
    cy.contains('Contract').click()
    cy.url().should('include', '/browse-contracts')
    //Had to do get with selector for this to work?
    //Look at all three tables (messages then notifications then contracts)
    cy.get('body > app-root > app-browse-contracts > div > div > div:nth-child(1) > mat-sidenav-container > mat-sidenav-content > mat-nav-list > a:nth-child(2) > span').click()
    cy.contains('Recipients')
    cy.get('body > app-root > app-browse-contracts > div > div > div:nth-child(1) > mat-sidenav-container > mat-sidenav-content > mat-nav-list > a:nth-child(1) > span').click()
    cy.contains('Required Action')
    cy.get('body > app-root > app-browse-contracts > div > div > div:nth-child(1) > mat-sidenav-container > mat-sidenav-content > mat-nav-list > a:nth-child(3) > span').click()
    cy.contains('Contract ID')
  })
})


describe('Loads the Contracts', () => {
  it('Loads the Contracts', () => {
    cy.viewport(1024, 768);
    //Has the example contract listed on the click of the button (would show fake contract here otherwise).
    cy.contains('example contract');
  
  })
})

describe('Finish Demo', () => {
  it('Finishes the Demo', () => {
  //Need this so the video doesn't cut off.
    cy.wait(1000);
  })
})