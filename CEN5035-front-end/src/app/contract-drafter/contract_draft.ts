export class ContractDraft {
    constructor(
      public UserID?: string,
      public AttorneyList?: string,
      public ClientList?:   string,
      public ContractTitle?: string,
      public Date?:          string,   
      public TerminationDate?: string,
      public PaymentType?: string,
      public OtherNotes?:  string
    ) {}
  }