package parsers

type computerFactory struct {
    programBuilder ProgramBuilder
    executionsBuilder ExecutionsBuilder
    parametersBuilder ParametersBuilder
    parameterBuilder ParameterBuilder
    executionBuilder ExecutionBuilder
    applicationBuilder ApplicationBuilder
    variablesBuilder VariablesBuilder
    variableBuilder VariableBuilder
}

func createComputerFactory(
    programBuilder ProgramBuilder,
    executionsBuilder ExecutionsBuilder,
    parametersBuilder ParametersBuilder,
    parameterBuilder ParameterBuilder,
    executionBuilder ExecutionBuilder,
    applicationBuilder ApplicationBuilder,
    variablesBuilder VariablesBuilder,
    variableBuilder VariableBuilder,
    ) ComputerFactory {
    out := computerFactory{
        programBuilder: programBuilder,
        executionsBuilder:executionsBuilder,
        parametersBuilder:parametersBuilder,
        parameterBuilder: parameterBuilder,
        executionBuilder: executionBuilder,
        applicationBuilder: applicationBuilder,
        variablesBuilder: variablesBuilder,
        variableBuilder: variableBuilder,
    }
    return &out
}

// Create creates a new computer instance
func (app *computerFactory) Create() Computer {
    return createComputer(
        app.programBuilder,
        app.executionsBuilder,
        app.parametersBuilder,
        app.parameterBuilder,
        app.executionBuilder,
        app.applicationBuilder,
        app.variablesBuilder,
        app.variableBuilder,
    )
}
