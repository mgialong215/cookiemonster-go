package src
import fs
func main(){
	type Wallet struct{
		network: string
		adress: string
		key: object 
		name: string 
	}
	type RewardSource struct{
	// 	type: 'liquiditypool' | 'staking'
    // address?: string,
    // api: string,
	}
	function loadWallets(root:string): Wallet[] {
		fs.readFileSync(root);
		// Go to directory, iterate through key files.
     	// Use filename to map to network
     	// network-keyname.json
		return []
		function loadRewardSources(wallet:Wallet)
}
os.Setenv(loadWallets,Wallet)  