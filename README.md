# SVM
Steve's Virtual Machine (SVM) is a simplistic virtual machine built to be easily parsed, extended and embedded into other software.

### Module
#### Declaration
```
// module declaration:
module myModule;
```

#### Type Declaration
##### Data type declaration
```
// type declaration:
type data myModule.myDataType;
```

##### Application type declaration
```
// type declaration:
type application myModule.myApplicationType;
```

#### Variable Declaration
```
// variable declaration:
myModule.myType $myVariable;
```

#### Parameter declaration
The input parameters can be of any data type.
```
// this is variable with an input:
-> myModule.myDataType $myInput;

// this is a variable with an output:
<- myModule.myDataType $myOutput;
```

#### Variable Assignment
```
// declaration with assignment:
myModule.myType $myVariable = ANY VALUE EXCEPT NON-ESCAPED SEMI-COLON;

// assignment of a previously declared variable:
myVariable = ANY VALUE EXCEPT NON-ESCAPED SEMI-COLON;
```

#### Attachment (attach/detach)
```
// attach variable to application:
attach myDataVariable:data @ myAppVariable;

// detach variable from application:
detach myDataVariable:data @ myAppVariable;
```
### Execution
```
// execute module application that returns a new variable:
myModule.myDataType $myOutput = execute myAppVariable;

// execute module application that returns an already declared variable:
$myOutput = execute myAppVariable;

// execute module application without a return clause:
execute myAppVariable;
```
