
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
    .type('fakeaccount@fakeaccount.com')
    .should('have.value','fakeaccount@fakeaccount.com')
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
    .type('fakeaccount@fakeaccount.com')
    .should('have.value','fakeaccount@fakeaccount.com')
    cy.get('input[name="password"]')
    .type('password1')
    .should('have.value','password1')
    })
})

describe('Opens the dashboard', () => {
  it('Open the user dashboard', () => {
    //Without this there are some button visibility issues?
    cy.viewport(1024, 768)
    cy.get('body > app-root > app-login-form > section > div > div > div.columns.is-centered > div > form > div.field.is-grouped > div:nth-child(1) > button').click()
    cy.url().should('include', '/users')
    cy.contains('Dashboard')
  })
})

describe('Opens page to send message', () => {
  it('Opens the message draft form', () => {
    //Without this there are some button visibility issues?
    cy.viewport(1024, 768)
    cy.contains('Send Messages').click()
    cy.url().should('include', '/message-draft')
    cy.get('body > app-root > app-message-drafter > form > div:nth-child(2) > div > div > input').type('fakeaccount@fakeaccount.com')
    cy.get('body > app-root > app-message-drafter > form > div:nth-child(3) > div > div > textarea').type('This is the test message')
    cy.get('body > app-root > app-message-drafter > form > div.field.is-grouped > div:nth-child(1) > div > button').click()

  })
})

describe('Opens contracts page and correctly interacts with form', () => {
  it('Opens contracts page and correctly interacts with form', () => {
    //Without this there are some button visibility issues?
    cy.viewport(1024, 768)
    cy.get('body > app-root > app-user-list > div > div > div.column.is-9 > div > div:nth-child(1) > app-field-list:nth-child(1) > div > footer > a').click()
    cy.url().should('include', '/contract-draft')
    cy.contains('Send a Contract')
    cy.get('body > app-root > app-contract-drafter > form > div:nth-child(2) > div > div > div > div.control.is-expanded > input').type('FakeClientEmail@email.com')
    cy.get('body > app-root > app-contract-drafter > form > div:nth-child(3) > div > div > div > div.control.is-expanded > input').type('Fake Client')
    cy.get('body > app-root > app-contract-drafter > form > div:nth-child(4) > div > div > input').type('Imaginary')
    cy.get('body > app-root > app-contract-drafter > form > div:nth-child(6) > div > div > input').type('Rocks')
    cy.get('body > app-root > app-contract-drafter > form > div:nth-child(7) > div > div > input').type('0')
    cy.get('body > app-root > app-contract-drafter > form > div:nth-child(8) > div > div > input').type('5')
    //Upload the file now.
    cy.get('body > app-root > app-contract-drafter > form > div:nth-child(9) > div > label > input').selectFile('cypress/fixtures/sample.pdf')
    cy.get('body > app-root > app-contract-drafter > form > div.field.is-grouped > div:nth-child(3) > div > button').click()
  })
})

describe('Searches for messages', () => {
  it('Searches for messages', () => {
    //Without this there are some button visibility issues?
    cy.viewport(1024, 768)
    //Upload the file now.
    cy.get('body > app-root > app-user-list > div > div > div.column.is-9 > div > div:nth-child(1) > app-field-list:nth-child(2) > div > div.card.my-5 > div > div > div > div > input').type('fakeaccount@fakeaccount.com')
    cy.get('body > app-root > app-user-list > div > div > div.column.is-9 > div > div:nth-child(1) > app-field-list:nth-child(2) > div > div.card-table > div > div > table > tbody > tr:nth-child(2) > td:nth-child(2)').contains('fakeaccount@fakeaccount.com')
  })
})

describe('Searches for contracts', () => {
  it('Searches for contracts', () => {
    //Without this there are some button visibility issues?
    cy.viewport(1024, 768)
    //Upload the file now.
    cy.get('body > app-root > app-user-list > div > div > div.column.is-9 > div > div:nth-child(1) > app-field-list:nth-child(1) > div > div.card.my-5 > div > div > div > div > input').type('00000001')
    cy.get('body > app-root > app-user-list > div > div > div.column.is-9 > div > div:nth-child(1) > app-field-list:nth-child(1) > div > div.card-table > div > div > table > tbody > tr:nth-child(2) > td:nth-child(3)').contains('Fake Client')
  })
})

describe('View a Message', () => {
  it('Views a Message', () => {
    //Without this there are some button visibility issues?
    cy.viewport(1024, 768)
    //Upload the file now.
    cy.get('body > app-root > app-user-list > div > div > div.column.is-9 > div > div:nth-child(1) > app-field-list:nth-child(2) > div > div.card-table > div > div > table > tbody > tr:nth-child(2) > td.level-right > a').click()
    cy.contains('This is the test message')
    cy.get('body > app-root > app-user-list > div > app-field-list-modal > div > div.modal-card > footer > button').click()
  })
})




describe('View a Contract', () => {
  it('Views a Contract', () => {
    //Without this there are some button visibility issues?
    cy.viewport(1024, 768)
    cy.get('body > app-root > app-user-list > div > div > div.column.is-9 > div > div:nth-child(1) > app-field-list:nth-child(1) > div > div.card-table > div > div > table > tbody > tr:nth-child(2) > td.level-right > a').click()
    cy.wait(100)
    //NOTE: Needed to be really basic with this because cypress has issues with PDF viewing: https://github.com/cypress-io/cypress/issues/2835
    cy.url().should('eq','http://localhost:4200/api/download?attorney_email=fakeaccount@fakeaccount.com&contract_id=00000001')
  })
})

describe('Finish Demo', () => {
  it('Finishes the Demo', () => {
  //Need this so the video doesn't cut off.
  cy.wait(1000);
})
})