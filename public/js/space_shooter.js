var gameboard = document.getElementById("gameboard");
var game = new Phaser.Game(gameboard.clientWidth, gameboard.clientHeight, Phaser.AUTO, "gameboard");

var bullets = new function() {
	this.types = [
		{
			damage : 2,
			title: 'SB 1',
			speed: 450,
			fireRate: 175,
			frameNum : 2,
			frame : { x : 6, y : 1, w : 15, h : 5}
		},
		{
			damage : 5,
			title: 'SB 2',
			speed: 450,
			fireRate: 150,
			frameNum : 0,
			frame : { x : 6, y : 1, w : 15, h : 5}
		},
		{
			damage : 10,
			title: 'HPB 1',
			speed: 450,
			fireRate: 115,
			frameNum : 3,
			frame : { x : 4, y : 1, w : 20, h : 6}
		},
		{
			damage : 15,
			title: 'HPB 2',
			speed: 450,
			fireRate: 100,
			frameNum : 1,
			frame : { x : 4, y : 1, w : 20, h : 6}
		},
		{
			damage : 25,
			title: 'Wave Cannon',
			speed: 400,
			fireRate : 120,
			frameNum : 8,
			frame : { x : 1, y : 0, w : 28, h : 10}
		},
		{
			damage : 25,
			title: 'HRF Laser Cannon',
			speed: 700,
			fireRate: 75,
			frameNum : 4,
			frame : { x : 0, y : 3, w : 30, h : 4}
		},
		{
			damage : 40,
			title: 'Plasma Disruptor',
			speed: 400,
			fireRate: 115,
			frameNum : 5,
			frame : { x : 10, y : 0, w : 10, h : 10}
		},
		{
			damage : 80,
			title: 'Plasma Cannon',
			speed: 400,
			fireRate: 120,
			frameNum : 6,
			frame : { x : 10, y : 0, w : 10, h : 10}
		},
		{
			damage : 160,
			title: 'Cosmic Disruptor',
			speed: 400,
			fireRate : 190,
			frameNum : 7,
			frame : { x : 20, y : 1, w : 10, h : 8}
		}
	];
	
	this.preload = function() {
		game.load.spritesheet('bullets', '/public/img/space_shooter/bullets.png', 30, 10);
	};
	
	this.create = function() {
		this.group = game.add.group();
		this.group.enableBody = true;
		this.group.createMultiple(150, 'bullets');
	};
	
	this.getFireRate = function(type) {
		return this.types[type].fireRate;
	};
	
	this.emitOne = function(x, y, type, angle) {
		var bullet = this.group.getFirstDead();
		
		bullet.reset(x, y);
		bullet.anchor.setTo(0, 0.5);
		if (angle) {
			bullet.rotation = angle * Math.PI / 180;
			game.physics.arcade.velocityFromAngle(angle, this.types[type].speed, bullet.body.velocity);
		} else {
			bullet.rotation = 0;
			bullet.body.velocity.x = this.types[type].speed;
		}
		bullet.body.offset.x = this.types[type].frame.x;
		bullet.body.offset.y = this.types[type].frame.y;
		bullet.body.width = this.types[type].frame.w;
		bullet.body.height = this.types[type].frame.h;
		bullet.checkWorldBounds = true;
		bullet.outOfBoundsKill = true;
		bullet.animations.add('p');
		bullet.type = this.types[type];
		bullet.animations.frame = this.types[type].frameNum;
	};
};

var spaceship = new function() {
	this.speed = 250;
	this.health = 10;
	this.fuel = 150;
	this.scale = 0.85;
	this.bulletType = 0;
	this.bulletSpread = 1;
	this.bulletTime = 0;
	this.fuelTime = 0;
	
	this.preload = function() {
		game.load.spritesheet('ship', '/public/img/space_shooter/spaceship.png', 100, 100, 9);
	},
	
	this.create = function() {
		this.ship = game.add.sprite(150, game.world.height / 2, 'ship');
		this.ship.animations.add('fly', [0, 3], true);
		this.ship.animations.add('bankRight', [1, 4], true);
		this.ship.animations.add('bankLeft',[2, 5], true);
		this.ship.animations.add('explode', [6,7,8], false);
		this.ship.scale.x = this.ship.scale.y = this.scale;
		
		this.ship.animations.play('fly', 15, true);
		game.physics.arcade.enable(this.ship);
		this.ship.body.collideWorldBounds = true;
		this.ship.body.offset.x = 0 * this.scale;
		this.ship.body.width = 42 * this.scale;
		this.ship.rotation = 90 * Math.PI / 180;
		this.ship.anchor.setTo(0.5, 0.5);
		
		this.healthbar = game.add.graphics(0, 0);
		this.drawnHealth = -100;
		this.healthLabel = game.add.text(game.world.width - 440, game.world.height - 20, "Health:", {
			font: '10px Arial',
			fill: '#00ff00'
		});
		this.fuelLabel = game.add.text(game.world.width - 190, game.world.height - 20, "Fuel:", {
			font: '10px Arial',
			fill: '#00ffff'
		});
		this.gunLabel = game.add.text(10, game.world.height - 20, bullets.types[this.bulletType].title, {
			font: '10px Arial',
			fill: '#00ff00'
		});
		
		this.fuelbar = game.add.graphics(0,0);
	};
	
	this.update = function(state) {
		this.ship.body.velocity.x = 0;
		this.ship.body.velocity.y = 0;
		
		this.burnFuel();
		
		if (this.exploding) {
			if (!this.exploding.isPlaying) {
				this.ship.kill();
				game.paused = true;
				return;
			}
		} else if (this.health <= 0) {
			this.exploding = this.ship.animations.play('explode', 9, false);
			return;
		} else {
			if (state.keyFire) {
				this.fireBullet();
			}
			if (state.keyLeft) {
				this.ship.body.velocity.x = -this.speed; 
			} else if (state.keyRight) {
				this.ship.body.velocity.x = this.speed;
			}
			if (state.keyUp) {
				this.ship.body.velocity.y = -this.speed;
				this.ship.animations.play('bankLeft', 15, true);
			} else if (state.keyDown) {
				this.ship.body.velocity.y = this.speed;
				this.ship.animations.play('bankRight', 15, true);
			} else {
				this.ship.animations.play('fly', 15, true);
			}
		}
	};
	
	this.reset = function() {
		this.bulletType = 1;
		this.bulletSpread = 1;
		this.health = 150;
	};
	
	this.incrBulletSpread = function() {
		this.bulletSpread += 1;
		if (this.bulletSpread > 5) { this.bulletSpread = 5; } 
		this.updateGunLabel();
	};
	
	this.incrBulletType = function() {
		if (this.bulletType < 8) {
			this.bulletType += 1;
			this.bulletSpread = 1;
			this.updateGunLabel();
		}
	};
	
	this.updateGunLabel = function() {
		var label = bullets.types[this.bulletType].title;
		if (this.bulletSpread > 1) {
			label = this.bulletSpread + " x " + label;
		}
		this.gunLabel.text = label;
	};
	
	this.incrFuel = function() {
		this.fuel += 25;
		if (this.fuel > 150) { this.fuel = 150; }
		this.drawFuelbar();
	};
	
	this.incrHealth = function() {
		this.health += 20;
		if (this.health > 150) { this.health = 150; }
	};
	
	this.drawHealthbar = function() {
		var color = 0x00FF00;
		if (this.health !== this.drawnHealth) {
			this.healthbar.clear();
			if (this.health < 85) {
				color = 0xFFFF00;
			}
			if (this.health < 40) {
				color = 0xFF0000;
			}
			this.healthbar.lineStyle(2, 0x00FF00, 1);
			this.healthbar.drawRect(game.world.width - 400, game.world.height - 20, 150, 10);
			this.healthbar.lineStyle(0, color, 0);
			this.healthbar.beginFill(color, 1);
			this.healthbar.drawRect(game.world.width - 400, game.world.height - 19, this.health-2, 8);
		}
	};
	
	this.drawFuelbar = function() {
		var color = 0x00FFFF;
		this.fuelbar.clear();
		this.fuelbar.lineStyle(2, color, 1);
		this.fuelbar.drawRect(game.world.width - 160, game.world.height - 20, 150, 10);
		this.fuelbar.beginFill(color, 1);
		this.fuelbar.drawRect(game.world.width - 160, game.world.height - 20, this.fuel, 10);
	};
	
	this.burnFuel = function() {
		if (game.time.now > this.fuelTime && this.fuel > 0) {
			this.fuel -= 1;
			this.drawFuelbar();
			
			this.fuelTime = game.time.now + 1000;
		}
	};
	
	this.fireBullet = function() {
		if (game.time.now > this.bulletTime) {
			var x = this.ship.body.position.x + 15;
			var y = this.ship.body.position.y;
			
			y += (this.ship.body.height / 2);
			
			if (this.bulletSpread === 1) {
				bullets.emitOne(x + 8, y, this.bulletType);
			} else if (this.bulletSpread === 2) {
				bullets.emitOne(x, y - 15, this.bulletType);
				bullets.emitOne(x, y + 15, this.bulletType);
			} else if (this.bulletSpread === 3) {
				bullets.emitOne(x, y - 20, this.bulletType, -15);
				bullets.emitOne(x, y + 20, this.bulletType, 15);
				bullets.emitOne(x + 8, y, this.bulletType, 0);
			} else if (this.bulletSpread === 4) {
				bullets.emitOne(x, y - 10, this.bulletType, -1.5);
				bullets.emitOne(x, y + 10, this.bulletType, 1.5);
				bullets.emitOne(x, y - 20, this.bulletType, -7);
				bullets.emitOne(x, y + 20, this.bulletType, 7);
			} else if (this.bulletSpread === 5) {
				bullets.emitOne(x+4, y - 15, this.bulletType, -3);
				bullets.emitOne(x+4, y + 15, this.bulletType, 3);
				bullets.emitOne(x, y - 25, this.bulletType, -8);
				bullets.emitOne(x, y + 25, this.bulletType, 8);
				bullets.emitOne(x + 8, y, this.bulletType, 0);
			}
			this.bulletTime = game.time.now + bullets.getFireRate(this.bulletType);
		}
	};
};

var mainState = {
	preload : function() {
		game.stage.backgroundColor = '#000000';
		
		game.load.image('stars', '/public/img/space_shooter/background.png');
		game.load.spritesheet('bonuses', '/public/img/space_shooter/bonus.png', 30, 30);
		bullets.preload();
		spaceship.preload();
	},
	
	create : function() {
		var state = this;
		game.physics.startSystem(Phaser.Physics.ARCADE);
		
		this.space = game.add.tileSprite(0, 0, game.world.width, game.world.height, 'stars');
		this.bonusGroup = game.add.group();
		this.bonusGroup.enableBody = true;
		this.bonusGroup.createMultiple(5, 'bonuses');
		
		bullets.create();
		spaceship.create();
		
		// setup game controlls
		this.leftKey = this.game.input.keyboard.addKey(Phaser.Keyboard.LEFT);
		this.rightKey = this.game.input.keyboard.addKey(Phaser.Keyboard.RIGHT);
		this.downKey = this.game.input.keyboard.addKey(Phaser.Keyboard.DOWN);
		this.upKey = this.game.input.keyboard.addKey(Phaser.Keyboard.UP);
		this.aKey = this.game.input.keyboard.addKey(Phaser.Keyboard.A);
		this.dKey = this.game.input.keyboard.addKey(Phaser.Keyboard.D);
		this.wKey = this.game.input.keyboard.addKey(Phaser.Keyboard.W);
		this.sKey = this.game.input.keyboard.addKey(Phaser.Keyboard.S);
		this.spacebar = this.game.input.keyboard.addKey(Phaser.Keyboard.SPACEBAR);
		
		this.bonusTimer = game.time.events.loop(2000, this.emitBonus, this);
	},
	
	update: function() {
		var state = {
			keyLeft : this.leftKey.isDown || this.aKey.isDown,
			keyRight : this.rightKey.isDown || this.dKey.isDown,
			keyUp : this.upKey.isDown || this.wKey.isDown,
			keyDown: this.downKey.isDown || this.sKey.isDown,
			keyFire: this.spacebar.isDown
		};
		
		this.space.tilePosition.x -= 2;
		spaceship.update(state);
		spaceship.drawHealthbar();
		
		game.physics.arcade.overlap(spaceship.ship, this.bonusGroup, this.addBonus, null, this);
	},
	
	emitBonus : function() {
		var t = Math.floor(Math.random() * 8);
		
		if (t < 4) {
			var bonus = this.bonusGroup.getFirstDead();
		
			var x = game.world.width;
			var y = Math.floor(Math.random() * (game.world.height - 50))+10;
		
			bonus.reset(x, y);
			bonus.body.velocity.x  = -200;
			bonus.checkWorldBounds = true;
			bonus.outOfBoundsKill  = true;
			bonus.animations.add('p');
			bonus.animations.frame = t;
			bonus.bonusType = t;
		}		
	},
	
	addBonus : function(ship, bonus) {
		if (bonus.bonusType === 0) {
			spaceship.incrHealth();
		} else if (bonus.bonusType === 1) {
			spaceship.incrBulletSpread();
		} else if (bonus.bonusType === 2) {
			spaceship.incrBulletType();
		} else {
			spaceship.incrFuel();
		}
		bonus.kill();
	}
};

game.state.add('main', mainState);  
game.state.start('main'); 