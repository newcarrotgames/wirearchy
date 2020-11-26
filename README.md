<p align="center">
  <img src="wirearchy.jpg">
</p>

### My latest over-ambitious project to create a game where all the content is generated using evolutionary algorithms.

So far, the code can take a structure like this:

```javascript
{
	"BaseName": "",
	"Created": "2020-11-14T10:46:18.8482241-05:00",
	"Nodes": [
		{
			"W": 0.4,
			"H": 0.4,
			"D": 0.4,
			"X": 0,
			"Y": 0,
			"Z": 0,
			"Material": 2,
			"Nodes": [
				{
					"W": 0.8,
					"H": 0.8,
					"D": 0.8,
					"X": 1.2,
					"Y": -1.5,
					"Z": 0.9,
					"Material": 3
				},
				{
					"W": 0.9,
					"H": 0.9,
					"D": 0.9,
					"X": -1.2,
					"Y": 1.5,
					"Z": 0.9,
					"Material": 4
				},
				{
					"W": 0.7,
					"H": 0.7,
					"D": 0.7,
					"X": 1.2,
					"Y": -1.5,
					"Z": -0.9,
					"Material": 6
				}
			]
		}
	]
}
```

and produce a voxel object like so:

<p align="center">
  <img src="update1.png">
</p>

Hopefully this system of components can model most of the content in the game. So far I've only really considered structures for POI generation but will eventually need to modify this framework to generate mobs.

## Creating a crude GAN generator/discrimator

Can create building-ish structures now using a random generator:

<p align="center">
  <img src="update2.png">
</p>

Working on a discriminator now.