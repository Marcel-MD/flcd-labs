## Abstract Syntax Tree

```
{
  type: PROGRAM
  statements: [
    {
      type: FUNCTION FLOAT
      identifier: {
        type: IDENTIFIER
        value: pow
      }
      parameters: [
        {
          type: PARAM FLOAT
          identifier: {
            type: IDENTIFIER
            value: a
          }
        }
        {
          type: PARAM INT
          identifier: {
            type: IDENTIFIER
            value: n
          }
        }
      ]
      body: {
        type: BLOCK
        statements: [
          {
            type: DECLARATION FLOAT
            identifier: {
              type: IDENTIFIER
              value: p
            }
            value: {
              type: IDENTIFIER
              value: a
            }
          }
          {
            type: FOR
            declaration: {
              type: DECLARATION INT
              identifier: {
                type: IDENTIFIER
                value: i
              }
              value: {
                type: INT_LITERAL
                value: 0
              }
            }
            condition: {
              type: INFIX LESS
              left: {
                type: IDENTIFIER
                value: i
              }
              operator: <
              right: {
                type: IDENTIFIER
                value: n
              }
            }
            increment: {
              type: PREFIX INCREMENT
              operator: ++
              right: {
                type: IDENTIFIER
                value: i
              }
            }
            body: {
              type: BLOCK
              statements: [
                {
                  type: ASSIGN
                  identifier: {
                    type: IDENTIFIER
                    value: p
                  }
                  operator: =
                  value: {
                    type: INFIX ASTERISK
                    left: {
                      type: IDENTIFIER
                      value: p
                    }
                    operator: *
                    right: {
                      type: IDENTIFIER
                      value: a
                    }
                  }
                }
              ]
            }
          }
          {
            type: RETURN
            value: {
              type: IDENTIFIER
              value: p
            }
          }
        ]
      }
    }
    {
      type: FUNCTION INT
      identifier: {
        type: IDENTIFIER
        value: main
      }
      parameters: [
      ]
      body: {
        type: BLOCK
        statements: [
          {
            type: DECLARATION FLOAT
            identifier: {
              type: IDENTIFIER
              value: p
            }
            value: {
              type: CALL
              function: {
                type: IDENTIFIER
                value: pow
              }
              arguments: [
                {
                  type: FLOAT_LITERAL
                  value: 2.13
                }
                {
                  type: INT_LITERAL
                  value: 10
                }
              ]
            }
          }
          {
            type: DECLARATION STR
            identifier: {
              type: IDENTIFIER
              value: s
            }
            value: {
              type: STR_LITERAL
              value:
            }
          }
          {
            type: IF
            condition: {
              type: INFIX AND
              left: {
                type: INFIX GREAT_EQ
                left: {
                  type: IDENTIFIER
                  value: p
                }
                operator: >=
                right: {
                  type: INT_LITERAL
                  value: 805
                }
              }
              operator: &&
              right: {
                type: INFIX NOT_EQ
                left: {
                  type: TRUE
                  value: true
                }
                operator: !=
                right: {
                  type: FALSE
                  value: false
                }
              }
            }
            consequence: {
              type: BLOCK
              statements: [
                {
                  type: ASSIGN
                  identifier: {
                    type: IDENTIFIER
                    value: s
                  }
                  operator: =
                  value: {
                    type: STR_LITERAL
                    value: success
                  }
                }
              ]
            }
            alternative: {
              type: BLOCK
              statements: [
                {
                  type: ASSIGN
                  identifier: {
                    type: IDENTIFIER
                    value: s
                  }
                  operator: =
                  value: {
                    type: STR_LITERAL
                    value: loss
                  }
                }
              ]
            }
          }
          {
            type: RETURN
            value: {
              type: INT_LITERAL
              value: 0
            }
          }
        ]
      }
    }
  ]
}
```
