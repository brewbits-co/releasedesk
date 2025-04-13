type ErrorResponse = {
  // A brief, user-friendly description of the error.
  message: string;
  // An array of strings providing additional context or suggestions to help resolve the error.
  helpTexts: string[];
};

type Channel = {
  ID: number;
  Name: string;
}