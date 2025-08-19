# RefleXSys 🤖📡 📶
**Self-Healing System Monitoring Framework *Inspired from Human Reflexes***

![Go](https://img.shields.io/badge/Go-1.24-blue)
[![CI/CD Checks](https://github.com/Coder-Harshit/RefleXSys/actions/workflows/go.yml/badge.svg)](https://github.com/Coder-Harshit/RefleXSys/actions/workflows/go.yml)
![License](https://img.shields.io/badge/License-MIT-green)
## Why RefleXSys ?
Most monitoring tools *react*.

RefleXSys is designed to **Anticipate, Decide & act**. 
This time it isn't just *graphs & alerts*,
rather 
it's about **building a nervous system for machines** - where agents watch, relay monitors & master thinks.

## Inspiration ?
*Why go outside when the the inspiration can be found in the very within ?*

RefleXSys derives its inspiration from the Reflex System of the Human Body.

Similar to how 'when we touch a **hot** item the signals travel within the nerve to the **spinal cord** and **brain** ; spinal cord instructs the muscle to move the hand back all the while also storing that the item is associated with something hot within the brain'.

## Architecture
![alt text](docs/Architecture.svg)

*Note:* Internals of the System maybe changed in future!

**Why? you may ask**.
> Because the Project is under Active Development 👷‍♂️.

### Components
**Agents**: The Agents are built with minimalism in mind & to be deployed on edge nodes. There whole task is to :
1. Monitor the resources of the system (it is deployed on)
2. Send(Report) the signals to the *Relay* node
3. Carry out the task received by the *Relay* Node

**Relay**: The Relay nodes maybe thought of as caching server in oversimplified terms. It is built with the vision that the agent would continuously sending the signals to it for monitoring purposes. Relay node would be responsible for:
1. Checking if the incoming signal is outside safety bound (need for supervision)
2. Taking action based on the ruleset recived from the master node (cached)
3. Grouping the incoming signal every *X* seconds and forwarding the batch to the master node 

**Master**: Would be acting as the Central Intelligence for the Framework. Resposiblities would include:
1. Deciding on what action to take
2. User Interactions (Connections to *OUTSIDE* world)
3. Keeping a check on Relay nodes
4. AddOn management
5. Analyzing the signal batch received!

**AddOns**: ***With Great Features comes Great Vulnerabilities!***, hence the vision of the Famework's Core components would be to provide the bare minimal and using the modular approach of AddOns to increment the functionalites.

> Each component would run independently, to make sure that if one fails the rest can continue working.


## Security
### Agents
Since the Agents would be deployed on real working machines, ensuring they work properly would take utmost priority. 

To make sure that the task is issued from the *authentic* Relay Node, unique ids would be assigned and hash matching shall be done before execution of any command.

Also for different level of operational sensitivity different level of log and access managment requirement shall be implemented.
*So that even if relay becomes faulty, the corrupted command won't be executed without proper access*

### Relay
Since the Relay need to monitor the signals from multiple agents so load balancing would be necessary (or at least something that make sures the relay isnt overburdened).
Also as the relay would storing the ruleset for what is normal and what is outside normal bounds, it need to update itself properly every *Y* seconds.

### Master
Since the master node is the one deciding on the commands and action to take & having even a single incorrect command issued can have serious impact, to avoid this, proper checks would be implemented before command is issued.

Level of automation would be INVERSELY PROPORTIONAL to the Sensitivity Involved

### AddOns
This area would mainly be the developer & security community to take charge for. Also it would be user's (or administrators) responsibility for the vulnerabilties and loopholes found within the AddOns they choose to use.
But not to worry, to lessen the amount of loopholes, verfied tags and ratings functionality shall be provided, so that users can choose what works best for them.


## What the Future Holds ?
The Framework would stay true to its promise of providing a nervous system for machines to **Anticipate, Decide & Act**.  


## 🚀 QuickStart Guide 

## Automated Build (via DOCKER-COMPOSE)

### 1. **Clone the Repo**

```bash
git clone https://github.com/Coder-Harshit/RefleXSys.git
cd RefleXSys
```

### 2. Build & Run the Entire System

*(Requires Docker and Docker Compose installed)*

```bash
docker compose up --build
```

- This will:
    - Build all (agent, relay, and master) images.
    - Start all services and connect them on a shared Docker network.
    - Mount the config files for usage in each service folder (`agent/config.yaml`, `master/config.yaml`, etc).
- **Logs from all services will appear in your terminal.**

### 3. Stopping the System

```bash
docker compose down
```

### 4. Configuration

- To customize port numbers, intervals, or endpoints, simply edit:
    - `agent/config.yaml`
    - `relay/config.yaml`
    - `master/config.yaml`
- Then restart with `docker compose up --build`.


## Manual Build/Run (via Go)

#### Building Binaries

```bash
cd agent  && go build -o build/agent main.go
cd ../relay && go build -o build/relay main.go
cd ../master && go build -o build/master main.go
```

#### Run Master Node (Default: **PORT 9101**)

```bash
cd master
go run main.go
```

#### Run Relay Node

```bash
cd relay
go run main.go
```

#### Run Agent Node

```bash
cd agent
go run main.go
```

***

### **Default Ports**

- **Master Node:** `9101`
- **Relay Node:** `8101`
- **Agent Node:** (does not listen, sends to relay)

Intervals, URLs, and ports are configurable via their respective `config.yaml` files.

## 🤝 Contribution
- Found a potential improvement ?
- Found a bug ?
  
>Make sure to raise *Issues* with proper tags!

- Have a BugFix ?

>Dont let that PR hold any longer, Raise it today!

<!-- Want to build a plugin? Open an issue. -->

