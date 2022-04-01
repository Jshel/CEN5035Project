

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
    })
})

describe('Opens the dashboard', () => {
  it('Open the user dashboard', () => {
    //Without this there are some button visibility issues?
    cy.viewport(1024, 768)
    cy.visit('/')
    cy.contains('Users').click()
    cy.url().should('include', '/users')
    cy.contains('Dashboard')
  })
})

describe('Opens and closes the Bulma modals', () => {
  it('Opens Bulma modals in the dashboard', () => {
    //Without this there are some button visibility issues?
    cy.viewport(1024, 768)
    cy.visit('/users')
    cy.get('body > app-root > app-user-list > div > div > div.column.is-9 > div > div:nth-child(1) > app-field-list:nth-child(1) > div > div:nth-child(2) > div > table > tbody > tr:nth-child(3) > td.level-right > a').click();
    cy.contains('Lorem ipsum')
    cy.get('body > app-root > app-user-list > div > app-field-list-modal > div > div.modal-card > header > button').click()
    //Check that modal can be escaped.
    cy.get('body > app-root > app-user-list > div > app-field-list-modal > div > div.modal-card > section').should('not.exist');
  })
})

describe('Opens messages page', () => {
  it('Opens the message draft form', () => {
    //Without this there are some button visibility issues?
    cy.viewport(1024, 768)
    cy.visit('/users')
    cy.get('body > app-root > app-user-list > div > div > div.column.is-9 > div > div:nth-child(1) > app-field-list:nth-child(2) > div > footer > a:nth-child(2)').click()
    cy.url().should('include', '/message-draft')
    cy.contains('Send a Message')
  })
})

describe('Opens notifications page', () => {
  it('Opens the notification draft form', () => {
    //Without this there are some button visibility issues?
    cy.viewport(1024, 768)
    cy.visit('/users')
    cy.get('body > app-root > app-user-list > div > div > div.column.is-9 > div > div:nth-child(1) > app-field-list:nth-child(3) > div > footer > a:nth-child(2)').click()
    cy.url().should('include', '/notification-draft')
    cy.contains('Create a Notification')
  })
})

describe('Opens contracts page and correctly interacts with form', () => {
  it('Opens contracts page and correctly interacts with form', () => {
    //Without this there are some button visibility issues?
    cy.viewport(1024, 768)
    cy.visit('/users')
    cy.get('body > app-root > app-user-list > div > div > div.column.is-9 > div > div:nth-child(1) > app-field-list:nth-child(1) > div > footer > a:nth-child(2)').click()
    cy.url().should('include', '/contract-draft')
    cy.contains('Send a Contract')
    //Test add and remove for attorneys.
    cy.get('body > app-root > app-contract-drafter > div:nth-child(3) > div > div > div > div.control.is-expanded > input').type('Fake name 1')
    cy.get('body > app-root > app-contract-drafter > div.field.is-grouped > div:nth-child(1) > div > button').click()
    cy.get('body > app-root > app-contract-drafter > div.field.is-grouped > div:nth-child(1) > div > button').click()
    cy.get('body > app-root > app-contract-drafter > div:nth-child(3) > div > div:nth-child(3) > div > div.control.is-expanded > input').type('Fake name 2')
    cy.get('body > app-root > app-contract-drafter > div:nth-child(3) > div > div:nth-child(4) > div > div.control.is-expanded > input').type('Fake name 3')
    //Remove second person and check that the element is gone.
    cy.get('body > app-root > app-contract-drafter > div:nth-child(3) > div > div:nth-child(3) > div > div:nth-child(2) > a').click()
    cy.contains('Fake name 2').should('not.exist')

    //Test add and remove for clients.
    cy.get('body > app-root > app-contract-drafter > div:nth-child(4) > div > div > div > div.control.is-expanded > input').type('Fake name 1')
    cy.get('body > app-root > app-contract-drafter > div.field.is-grouped > div:nth-child(2) > div > button').click()
    cy.get('body > app-root > app-contract-drafter > div.field.is-grouped > div:nth-child(2) > div > button').click()
    cy.get('body > app-root > app-contract-drafter > div:nth-child(4) > div > div:nth-child(3) > div > div.control.is-expanded > input').type('Fake name 2')
    cy.get('body > app-root > app-contract-drafter > div:nth-child(4) > div > div:nth-child(4) > div > div.control.is-expanded > input').type('Fake name 3')
    //Remove second person and check that the element is gone.
    cy.get('body > app-root > app-contract-drafter > div:nth-child(4) > div > div:nth-child(3) > div > div:nth-child(2) > a').click()
    cy.contains('Fake name 2').should('not.exist')

  })
})

describe('Finish Demo', () => {
  it('Finishes the Demo', () => {
  //Need this so the video doesn't cut off.
  cy.wait(1000);
})
})