// Package action defines actions for the simulation.
package action

// Action is a simulation input. Everything clients do that should provoke a change in the simulation
// should generate an Action.
// TODO: define a better interface. Details about the Action generation?
type Action interface{}
