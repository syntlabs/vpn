import UIKit

// Struct to hold the data needed for each ViewController
struct ViewControllerData {
    let descriptionText: String
    let backgroundColor: UIColor
    let buttonTitle: String
    let buttonAction: () -> Void
}

// AbstractViewController is a subclass of UIViewController and serves as a base class for other view controllers
class AbstractViewController: UIViewController {
    // IBOutlet for the UILabel that displays the description text
    @IBOutlet var descriptionLabel: UILabel!
    // IBOutlet for the UIButton that triggers the button action
    @IBOutlet var button: UIButton!
    // IBOutlet for the UIView that creates the parallax effect
    @IBOutlet var parallaxView: UIView!

    // IBInspectable property for the description text
    @IBInspectable var descriptionText: String?
    // IBInspectable property for the background color
    @IBInspectable var backgroundColor: UIColor?
    // IBInspectable property for the button title
    @IBInspectable var buttonTitle: String?
    // IBInspectable property for the button action
    @IBInspectable var buttonAction: String?
    // Property to keep track of the current index
    @IBInspectable var currentIndex: Int = 0

    override func viewDidLoad() {
        super.viewDidLoad()
        // Call the setupUI function to configure the user interface
        setupUI()
        // Call the updateUI function to update the user interface for the current index
        updateUI(forIndex: currentIndex)
    }

    // Function to setup the user interface
    func setupUI() {
        // Set the background color of the view to white
        view.backgroundColor = UIColor.white
        // Set the background color of the parallaxView to the background color property
        parallaxView.backgroundColor = backgroundColor
        // Set the text of the descriptionLabel to the descriptionText property
        descriptionLabel.text = descriptionText
        // Set the title of the button to the buttonTitle property
        button.setTitle(buttonTitle, for: .normal)
        // Add a target to the button that triggers the buttonTapped function when the button is tapped
        button.addTarget(self, action: #selector(buttonTapped), for: .touchUpInside)
    }

    // Function to update the user interface for the specified index
    func updateUI(forIndex index: Int) {
        // Set the current index to the specified index
        currentIndex = index
        // Get the ViewControllerData for the current index
        let viewControllerData = getViewControllerData(forIndex: index)
        // Set the descriptionText property to the descriptionText of the ViewControllerData
        descriptionText = viewControllerData.descriptionText
        // Set the backgroundColor property to the backgroundColor of the ViewControllerData
        backgroundColor = viewControllerData.backgroundColor
        // Set the buttonTitle property to the buttonTitle of the ViewControllerData
        buttonTitle = viewControllerData.buttonTitle
        // If the buttonAction property is not nil, remove the current target from the button and add a new target that triggers the specified action when the button is tapped
        if let actionString = buttonAction, let action = NSSelectorFromString(actionString) {
            button.removeTarget(nil, action: nil, for: .allEvents)
            button.addTarget(self, action: action, for: .touchUpInside)
        }
        // Call the setupUI function to configure the user interface with the updated properties
        setupUI()
    }

    // Function to get the ViewControllerData for the specified index
    func getViewControllerData(forIndex index: Int) -> ViewControllerData {
        // Array of ViewControllerData objects
        let viewControllerDataArray = [
            ViewControllerData(
                descriptionText: "This is View Controller 1",
                backgroundColor: .red,
                buttonTitle: "Next",
                buttonAction: { [weak self] in self?.nextButtonTapped() }
            ),
            ViewControllerData(
                descriptionText: "This is View Controller 2",
                backgroundColor: .green,
                buttonTitle: "Next",
                buttonAction: { [weak self] in self?.nextButtonTapped() }
            ),
            ViewControllerData(
                descriptionText: "This is View Controller 3",
                backgroundColor: .blue,
                buttonTitle: "Next",
                buttonAction: { [weak self] in self?.nextButtonTapped() }
            ),
            ViewControllerData(
                descriptionText: "This is View Controller 4",
                backgroundColor: .yellow,
                buttonTitle: "Lets go!",
                buttonAction: { [weak self] in self?.skipToMainViewController() }
            )
        ]
        // Return the ViewControllerData object at the specified index
        return viewControllerDataArray[index]
    }

    // Function that is triggered when the button is tapped
    @objc func buttonTapped() {
        // If the buttonAction property is not nil, perform the specified action
        if let action = buttonAction, let selector = NSSelectorFromString(action) {
            perform(selector)
        }
    }

    // Function that is triggered when the "Next" button is tapped
    func nextButtonTapped() {
        // Calculate the index of the next ViewController
        let nextIndex = currentIndex + 1
        // If the next index is less than 4, create a new instance of AbstractViewController, update its user interface for the next index, and push it onto the navigation stack. Animate the transition to the new ViewController.
        if nextIndex < 4 {
            let nextViewController = AbstractViewController()
            nextViewController.updateUI(forIndex: nextIndex)
            navigationController?.pushViewController(nextViewController, animated: false)
            animateTransition(to: nextViewController)
        }
    }

    // Function to skip to the MainViewController
    func skipToMainViewController() {
        // Create a new instance of MainViewController and pop to it on the navigation stack
        let mainViewController = MainViewController()
        navigationController?.popToViewController(mainViewController, animated: true)
    }

    // Function to animate the transition to the specified ViewController
    func animateTransition(to viewController: UIViewController?) {
        // If the viewController is nil, return from the function
        guard let viewController = viewController else { return }
        // Get the width and height of the screen
        let screenWidth = UIScreen.main.bounds.width
        let screenHeight = UIScreen.main.bounds.height
        // Get the current view and the view of the specified ViewController
        let fromView = view
        let toView = viewController.view
        // Calculate the x and y offsets for the parallax effect
        let xOffset = view.bounds.width * 0.2
        let yOffset = view.bounds.height * 0.2
        // Set the duration of the animation
        let animationDuration = 0.3
        // Animate the transition using a cross dissolve effect. Move the current view and the parallaxView of the specified ViewController to their final positions.
        UIView.transition(with: fromView, duration: animationDuration, options: .transitionCrossDissolve, animations: {
            fromView.transform = CGAffineTransform(translationX: -xOffset, y: -yOffset)
            toView.transform = CGAffineTransform.identity
        }, completion: { _ in
            // If the specified ViewController is an instance of AbstractViewController, animate the parallaxView to its final position
            if let viewController = viewController as? AbstractViewController {
                viewController.parallaxView.transform = CGAffineTransform(translationX: -xOffset, y: -yOffset)
                UIView.animate(withDuration: animationDuration, delay: 0, options: .curveEaseInOut, animations: {
                    viewController.parallaxView.transform = CGAffineTransform.identity
                })
            }
        })
    }
}