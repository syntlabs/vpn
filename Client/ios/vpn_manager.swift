import CFNetwork
import NetworkExtension

/**
 * VPNManager class is responsible for managing the VPN connection.
 * It provides methods to start and stop the VPN tunnel.
 */
class VPNManager {
    /**
     * Shared instance of NEVPNManager, which is used to manage the VPN configuration.
     */
    let vpnManager = NEVPNManager.shared()
    
    /**
     * Data representation of the password, stored as a reference to be used for authentication.
     */
    let passwordRef: Data
    
    /**
     * Initializes the VPNManager instance with a password.
     * The password is converted to Data using UTF-8 encoding and stored as a reference.
     *
     * - parameter password: The password to be used for authentication.
     */
    init(password: String) {
        passwordRef = password.data(using:.utf8)!
    }
    
    /**
     * Starts the VPN tunnel by loading the VPN configuration from preferences,
     * setting up the protocol configuration, and saving the changes to preferences.
     */
    func startVPN() {
        /**
         * Creates a new instance of NEVPNProtocol, which represents the VPN protocol configuration.
         */
        let vpnProtocol = NEVPNProtocol()
        
        /**
         * Sets the provider bundle identifier, which is used to identify the VPN provider.
         */
        vpnProtocol.providerBundleIdentifier = "com.example.vpn"
        
        /**
         * Sets the server address, which is the URL of the VPN server.
         */
        vpnProtocol.serverAddress = "your-proxy-server-url.com"
        
        /**
         * Sets the username, which is used for authentication.
         */
        vpnProtocol.username = "username"
        
        /**
         * Sets the password reference, which is used for authentication.
         */
        vpnProtocol.passwordReference = passwordRef
        
        /**
         * Sets the authentication method to none, which means no additional authentication is required.
         */
        vpnProtocol.authenticationMethod =.none
        
        /**
         * Enables extended authentication, which allows for more advanced authentication mechanisms.
         */
        vpnProtocol.useExtendedAuthentication = true
        
        /**
         * Creates a new instance of NEVPNCrypto, which represents the encryption and integrity algorithms.
         */
        let chacha20Encryption = NEVPNCrypto()
        
        /**
         * Sets the encryption algorithm to ChaCha20, which is a fast and secure encryption algorithm.
         */
        chacha20Encryption.encryptionAlgorithm =.chacha20
        
        /**
         * Sets the integrity algorithm to Poly1305, which is a fast and secure integrity algorithm.
         */
        chacha20Encryption.integrityAlgorithm =.poly1305
        
        /**
         * Sets the crypto configuration to the VPN protocol.
         */
        vpnProtocol.crypto = chacha20Encryption
        
        /**
         * Loads the VPN configuration from preferences, which may contain previously saved settings.
         * If an error occurs, it is printed to the console.
         */
        vpnManager.loadFromPreferences { error in
            if let error = error {
                print("Error loading VPN preferences: \(error)")
                return
            }
            
            /**
             * Sets the protocol configuration to the VPN manager.
             */
            self.vpnManager.protocolConfiguration = vpnProtocol
            
            /**
             * Enables on-demand VPN, which allows the VPN to start automatically when needed.
             */
            self.vpnManager.isOnDemandEnabled = true
            
            /**
             * Saves the VPN configuration to preferences, which persists the changes.
             * If an error occurs, it is printed to the console.
             */
            self.vpnManager.saveToPreferences { error in
                if let error = error {
                    print("Error saving VPN preferences: \(error)")
                    return
                }
                
                /**
                 * Starts the VPN tunnel, which establishes the connection to the VPN server.
                 */
                self.vpnManager.startVPNTunnel()
            }
        }
    }
    
    /**
     * Stops the VPN tunnel, which disconnects the VPN connection.
     */
    func stopVPN() {
        /**
         * Stops the VPN tunnel, which disconnects the VPN connection.
         */
        vpnManager.stopVPNTunnel()
    }
}