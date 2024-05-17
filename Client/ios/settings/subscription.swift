import UIKit

class SubscriptionStatusViewController: UIViewController {

    @IBOutlet weak var countdownLabel: UILabel!

    let dateFormatter = DateFormatter()
    let subscriptionEndDate: Date = Date().addingTimeInterval(30 * 24 * 60 * 60) // 30 days from now

    override func viewDidLoad() {
        super.viewDidLoad()

        title = "Subscription Status"

        countdownLabel.text = "Subscription ends in: "
        updateCountdown()
    }

    func updateCountdown() {
        let timeRemaining = subscriptionEndDate.timeIntervalSinceNow
        let days = Int(timeRemaining / 86400)
        let hours = Int((timeRemaining % 86400) / 3600)
        let minutes = Int((timeRemaining % 3600) / 60)
        let seconds = Int(timeRemaining % 60)

        countdownLabel.text = "Subscription ends in: \(days) days, \(hours) hours, \(minutes) minutes, \(seconds) seconds"

        Timer.scheduledTimer(timeInterval: 1, target: self, selector: #selector(updateCountdown), userInfo: nil, repeats: true)
    }
}