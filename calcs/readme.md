# calcs module

This pkg is responsible for taking the data collected from
the `proc` pkg and doing the kind of esoteric transformations
needed to put them into the formats that `draw` needs.

Some of the transformed data will be in mutable chunks of 
the `state.State` struct.