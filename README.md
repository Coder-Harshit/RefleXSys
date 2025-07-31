# RefleXSys 🤖📡 📶
**Self-Healing System Monitoring Framework *Inspired from Human Reflexes***

![Go](https://img.shields.io/badge/Go-1.24-blue)
[![CI/CD Checks](https://github.com/Coder-Harshit/RefleXSys/actions/workflows/go.yml/badge.svg)](https://github.com/Coder-Harshit/RefleXSys/actions/workflows/go.yml)
![License](https://img.shields.io/badge/License-MIT-green)
## Why RefleXSys ?
Most monitoring tools act as *react*.

RefleXSys is designed to **Anticipate, Decide & act**. 
This time it isn't just *graphs & alerts*,
rather 
it's about **building a nervous system for machines** - where agents watch,master thinks & addons respond - AUTOMATICALLY.

## Inspiration ?
*Why go outside when the the inspiration can be found in the very within ?*

RefleXSys derives its inspiration from the Reflex System of the Human Body

## Architecture
![alt text](docs/Architecture.svg)

*Note:* This is a very simplified version of the System

**Why? you may ask**.
> Because the Project is under Active Development 👷‍♂️.

### Components
**Agents**: The Agents are built with minimalism in mind & to be deployed on edge nodes. There whole task is to :
1. Monitor the resources of the system (it is deployed on)
2. Send(Report) the signals to the *Master* node with varying *sensitivity rating*
3. Carry out the task received by the Master Node

**Master**: Would be acting as the Central Intelligence for the Framework. Resposiblities would include:
1. Deciding on what action to take
2. User Interactions (Connections to *OUTSIDE* world)
3. Keeping a check on Agents'

**AddOns**: ***With Great Features comes Great Vulnerabilities!***, hence the vision of the Famework's Core components would be to provide the bare minimal and using the modular approach of AddOns to increment the functionalites.

> Each component would run independently, to make sure that if one fails the rest can continue working.


## Security
### Agents
Since the Agents would be deployed on real working machines, ensuring they work properly would take utmost priority. 
To make sure that the task is issued from the *authentic* Master Node, unique ids would be assigned and hash matching shall be done before execution of any command.
Also for different level of operational sensitivity different level of log and access managment requirement shall be implemented.

### Master
Since the master node is the one issuing commands and responsible for deciding what action to take & having even a single incorrect command issued can have serious impact, to avoid this, proper checks would be implemented before command is issued and also sensitivity of the topic would be checked.

Level of automation would be INVERSELY PROPORTIONAL to the Sensitivity Involved

### AddOns
This area would mainly be the developer & security community to take charge for. Also it would be user's (or administrators) responsibility for the vulnerabilties and loopholes found within the AddOns they choose to use.
But not to worry, to lessen the amount of loopholes, verfied tags and ratings functionality shall be provided, so that users can choose what works best for them.


## What the Future Holds ?
The Framework would stay true to its promise of providing a nervous system for machines to **Anticipate, Decide & Act**.  

## QuickStart Guide
1. Repo Cloning
```bash
git clone https://github.com/Coder-Harshit/RefleXSys.git
cd RefleXSys
```
2. Building Agent
```bash
cd agent && go build -o build/agent main.go
```

3. Building Master
```bash
cd ../master && go build -o build/master main.go
```
4. Running Master Node
```bash
go run main.go
```

5. Running Agent Node
```bash
cd agent && go run main.go
```

> Master Node runs on **PORT 8080**
> 
> Agent Node sends signal every second

## 🤝 Contribution
- Found a potential improvement ?
- Found a bug ?
  
>Make sure to raise *Issues* with proper tags!

- Have a BugFix ?

>Dont let that PR hold any longer, Raise it today!

<!-- Want to build a plugin? Open an issue. -->

