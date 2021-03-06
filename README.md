# Go-Computational-Fluid-Dynamics
Learning Go by trying to implement a basic computational fluid dynamics simulation.

Following tutorial : http://nbviewer.ipython.org/github/barbagroup/CFDPython/tree/master/lessons/ 

## Current

Step 7 - 2D Diffusion

Go into folder step7_2D_Diffusion and call runTest.sh

	cd step7_2D_Diffusion
	chmod +x runTest.sh
	./runTest.sh

![Example sim](1d_simN400x1000S.png)

## Previously
Started with a 1D simulation of pressure diffusion. We take a zeroed 1D array of length N
with seeded intial pressures of an ideal gas, and run for a number of Steps.

This is visualized by drawing each step as a column in a N x Steps sized PNG image. Each column slice is one time step.
The color of the pixel is dictated by the pressure (say white is max pressure in that timestep), or based on overall diffusion.
