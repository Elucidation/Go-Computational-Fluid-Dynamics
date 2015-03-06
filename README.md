# Go-Computational-Fluid-Dynamics
Learning Go by trying to implement a basic computational fluid dynamics simulation

Starting with a 1D simulation of pressure diffusion. We take a zeroed 1D array of length N
with seeded intial pressures of an ideal gas, and run for a number of Steps.

![Example sim](https://github.com/Elucidation/Go-Computational-Fluid-Dynamics/blob/master/1d_simN400x400S.png)

This is visualized by drawing each step as a column in a N x Steps sized PNG image. Each column slice is one time step.

The color of the pixel is dictated by the pressure (say white is max pressure in that timestep), or based on overall diffusion.
