package main

func isProdEnv() bool {
	// balabala
	return true
}

func degradeDeploy() bool {
	// balabala
	return true
}

func isChaosAgent() bool {
	// balabala
	return true
}

func needAudingDeploy() bool {
	if !isProdEnv() {
		return false
	}

	if !degradeDeploy() {
		return false
	}

	if isChaosAgent() {
		return false
	}

	return true
}

func needAudingDeployV2() bool {
	if isProdEnv() && degradeDeploy() && !isChaosAgent() {
		return true
	}
	return false
}
