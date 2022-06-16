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
module type data myDataType;
```

##### Application type declaration
```
// type declaration:
module type application myApplicationType;
```

#### Variable Declaration
```
// data variable declaration:
myModule.myDataType myDataVariable;

// application variable declaration:
myModule.myApplicationType myAppVariable;
```

#### Parameters declaration
The input parameters can be of any data type.
```
// this isa variable with an input:
-> myModule.myDataType myInput;

// this is a variable with an output:
<- myModule.myDataType myOutput;
```

#### Variable Assignment
```
// declaration with assignment:
myModule.myDataType myDataVariable = ANY VALUE EXCEPT NON-ESCAPED SEMI-COLON;

// assignment of a previously declared variable:
myModule.myApplicationType myAppVariable = ANY VALUE EXCEPT NON-ESCAPED SEMI-COLON;
```

#### Attachment (attach/detach)
```
// attach bytes to application:
attach bytes myData:data to myModule.myApplicationType myAppVariable;

// attach module data variable to application:
attach myModule.myDataType myModule.myDataVariable:data to myModule.myApplicationType myAppVariable;

// detach bytes from application:
detach bytes myData:data to myModule.myApplicationType myAppVariable;

// detach module data variable from application:
detach myModule.myDataType myModule.myDataVariable:data from myModule.myApplicationType myAppVariable;
```
### Execution
```
// execute module application that returns a module output:
myModule.myDataType myModule.myOutput = execute myModule.myApplicationType myAppVariable;

// execute module application that return bytes:
bytes output = execute myModule.myApplicationType myAppVariable;

// execute module application without a return clause:
execute myModule.myApplicationType myAppVariable;
```
