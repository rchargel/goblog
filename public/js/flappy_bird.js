var gameboard = document.getElementById("gameboard");
var game = new Phaser.Game(gameboard.clientWidth, gameboard.clientHeight, Phaser.AUTO, "gameboard");

var fruitType = [
	{
		name : "mango",
		points : 30,
		frame : 0
	},
	{
		name : "banana",
		points : 10,
		frame : 2
	},
	{
		name : "walnut",
		points : 5,
		frame : 1
	}
]
var mainState = {
	preload: function() {
		this.flappingSpeed = { 
			fast: 25,
			slow: 12
		};
		this.speed = 250;
		this.birdVelocity = 300;
		this.gravity = 1000;
		this.yScale = 20;
		this.yPositions = (game.world.height - 70) / this.yScale;
		game.stage.backgroundColor = '#146d2c';
		game.load.image('jungle', '/public/img/flappy_bird/background.png');
		game.load.spritesheet('fruit', '/public/img/flappy_bird/fruit.png', 50, 50, 3);
		game.load.spritesheet('macaw', '/public/img/flappy_bird/macaws.png', 135, 100, 8);
		game.load.spritesheet('vines', '/public/img/flappy_bird/vine.png', 50, 487, 4);
	},
	
	create: function() {
		game.physics.startSystem(Phaser.Physics.ARCADE);
		
		// create background
		this.jungle = game.add.tileSprite(0, 0, 1000, 500, 'jungle');
		
		// create fruit
		this.fruit = game.add.group();
		this.fruit.enableBody = true;
		this.fruit.createMultiple(30, 'fruit');
		
		// craete vines
		this.vines = game.add.group();
		this.vines.enableBody = true;
		this.vines.createMultiple(30, 'vines');
		
		// create bird
		this.bird = this.game.add.sprite(100, (game.world.height / 2) - 37, 'macaw');
		this.bird.animations.add('fly');
		this.bird.animations.play('fly', this.flappingSpeed.slow, true);
		this.bird.scale.x = this.bird.scale.y = 1;
		
		// create score
		this.score = 0;
		this.scoreLabel = game.add.text(20, 20, "0", {
			font: '30px Arial',
			fill: '#ffffff'
		});
		
		// create timers
		this.emitTimer = game.time.events.loop(4000, this.emit, this);
		this.speedTime = game.time.events.loop(6000, this.incrementSpeed, this);
		
		// setup game controlls
		this.game.input.onDown.add(this.flyUp, this);
		this.game.input.onUp.add(this.flyDown, this);
		var spacebar = this.game.input.keyboard.addKey(Phaser.Keyboard.SPACEBAR);
		spacebar.onDown.add(this.flyUp, this);
		spacebar.onUp.add(this.flyDown, this);
		this.mouseDown = false;
		
		// setup physics
		game.physics.arcade.enable(this.bird);
    	this.bird.body.gravity.y = this.gravity;
		this.bird.body.collideWorldBounds = true;
		this.bird.body.offset.y = 47;
		this.bird.body.offset.x = 63;
		this.bird.body.width = 63;
		this.bird.body.height = 22;
		
		this.game.paused = true;
	},
	
	update: function() {
		if (this.mouseDown) {
			this.bird.body.velocity.y = -this.birdVelocity;
		} 
		this.jungle.tilePosition.x -= (this.speed / 100);
		game.physics.arcade.overlap(this.bird, this.vines, this.restartGame, null, this);
		game.physics.arcade.overlap(this.bird, this.fruit, this.incrementScore, null, this);
	},
	
	restartGame: function() {
		game.paused = true;
		game.state.start('main');
	},
	
	flyUp: function() {
		if (this.game.paused) {
			this.game.paused = false;
		} else {
			this.bird.animations._anims['fly'].speed = this.flappingSpeed.fast;
			this.mouseDown = true;
		}
	},
	
	flyDown: function() {
		this.bird.animations._anims['fly'].speed = this.flappingSpeed.slow;
		this.mouseDown = false;
	},
	
	randLocation: function() {
		return {
			position: Math.floor(Math.random() * 5),
			incr: game.world.height / 5,
			opening: (Math.floor(Math.random() * 2) + 2) * (game.world.height / 5)
		};
	},
	
	emit: function() {
		var location = this.randLocation();
		this.emitVines(location);
		this.emitFruit(location);
	},
	
	emitVines: function(loc) {
		var x = game.world.width;
		var openSize = loc.opening;
		var qty = Math.floor(Math.random() * 3) + 1;
		
		x += ((4 - qty) * 100) / 2
		
		for (var i = 0; i < qty; i++) {
			var openLoc = loc.position * loc.incr;
			if (openLoc > 0) {
				this.emitVine(x + (i * 100), openLoc - 487, false);
			}
			if ((openLoc + openSize) < game.world.height) {
				this.emitVine(x + (i * 100), openLoc + openSize, true);
			}
		}
	},
	
	emitVine: function(x, y, up) {
		var frame = Math.floor(Math.random() * 2);
		frame += (up ? 0 : 2);
		var vine = this.vines.getFirstDead();
		
		vine.reset(x, y);
		vine.body.velocity.x = -this.speed;
		vine.checkWorldBounds = true;
		vine.events.onOutOfBounds.add(function(f) {
			if (f.position.x < 0) {
				f.kill();
			}
		}, this);
		vine.animations.add('v');
		vine.animations.frame = frame;
	},
	
	emitFruit: function(loc) {
		var x = game.world.width;
		var y = (loc.position * loc.incr) + (loc.opening / 2 - 20);
		
		if (y < 5) { 
			y = 5;
		}
		
		if (y > (game.world.height - 50)) {
			y = game.world.height - 50;
		}
		
		var num = Math.floor(Math.random() * 5) + 4;
		
		for (var i = 0; i < num; i++) {
			this.emitOne(x + (i * 70), y);
		}
	},
	
	emitOne: function(x, y) {
		var oneFruit = this.fruit.getFirstDead();
		var r = (Math.floor(Math.random() * 8) - 4) / 10;
		var type = Math.floor(Math.random() * 100);
		if (type > 90) {
			type = 0;
		} else if (type > 50) {
			type = 1;
		} else {
			type = 2;
		}
		oneFruit.reset(x, y);
		oneFruit.body.velocity.x = -this.speed;
		oneFruit.body.offset.y = 3;
		oneFruit.body.offset.x = 8;
		oneFruit.body.width = 34;
		oneFruit.body.height = 44;
		oneFruit.rotation = r;
		oneFruit.scale.x = oneFruit.scale.y = 0.8;
		oneFruit.checkWorldBounds = true;
		oneFruit.events.onOutOfBounds.add(function(f) {
			if (f.position.x < 0) {
				f.kill();
			}
		}, this);
		//oneFruit.outOfBoundsKill = true;
		oneFruit.animations.add('p');
		oneFruit.type = fruitType[type];
		oneFruit.animations.frame = oneFruit.type.frame;
	},
	
	incrementSpeed: function() {
		if (this.emitTimer.delay > 500) {
			this.speed += (this.speed * 0.05);
			this.emitTimer.delay = this.emitTimer.delay * 0.95;
		}
	},
	
	incrementScore: function(bird, fruit) {
		var val = fruit.type.points;
		fruit.kill();
		this.score += val;
		this.scoreLabel.text = this.score;
	}
};
game.state.add('main', mainState);  
game.state.start('main'); 
