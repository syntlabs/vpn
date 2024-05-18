import UIKit

struct ViewControllerData {
    let descriptionText: String
    let backgroundColor: UIColor
    let buttonTitle: String
    let buttonAction: () -> Void
}

class AbstractViewController: UIViewController {
    @IBOutlet var descriptionLabel: UILabel!
    @IBOutlet var button: UIButton!
    @IBOutlet var parallaxView: UIView!

    @IBInspectable var descriptionText: String?
    @IBInspectable var backgroundColor: UIColor?
    @IBInspectable var buttonTitle: String?
    @IBInspectable var buttonAction: String?
    @IBInspectable var currentIndex: Int = 0

    override func viewDidLoad() {
        super.viewDidLoad()
        setupUI()
        updateUI(forIndex: currentIndex)
    }

    func setupUI() {
        view.backgroundColor = UIColor.white
        parallaxView.backgroundColor = backgroundColor
        descriptionLabel.text = descriptionText
        button.setTitle(buttonTitle, for: .normal)
        button.addTarget(self, action: #selector(buttonTapped), for: .touchUpInside)
    }

    func updateUI(forIndex index: Int) {
        currentIndex = index
        let viewControllerData = getViewControllerData(forIndex: index)
        descriptionText = viewControllerData.descriptionText
        backgroundColor = viewControllerData.backgroundColor
        buttonTitle = viewControllerData.buttonTitle
        if let actionString = buttonAction, let action = NSSelectorFromString(actionString) {
            button.removeTarget(nil, action: nil, for: .allEvents)
            button.addTarget(self, action: action, for: .touchUpInside)
        }
        setupUI()
    }

    func getViewControllerData(forIndex index: Int) -> ViewControllerData {
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
        return viewControllerDataArray[index]
    }

    @objc func buttonTapped() {
        if let action = buttonAction, let selector = NSSelectorFromString(action) {
            perform(selector)
        }
    }

    func nextButtonTapped() {
        let nextIndex = currentIndex + 1
        if nextIndex < 4 {
            let nextViewController = AbstractViewController()
            nextViewController.updateUI(forIndex: nextIndex)
            navigationController?.pushViewController(nextViewController, animated: false)
            animateTransition(to: nextViewController)
        }
    }

    func skipToMainViewController() {
        let mainViewController = MainViewController() 
        navigationController?.popToViewController(mainViewController, animated: true)
    }

    func animateTransition(to viewController: UIViewController?) {
        guard let viewController = viewController else { return }
        let screenWidth = UIScreen.main.bounds.width
        let screenHeight = UIScreen.main.bounds.height

        let fromView = view
        let toView = viewController.view

        let xOffset = view.bounds.width * 0.2
        let yOffset = view.bounds.height * 0.2
        let animationDuration = 0.3

        UIView.transition(with: fromView, duration: animationDuration, options: .transitionCrossDissolve, animations: {
            fromView.transform = CGAffineTransform(translationX: -xOffset, y: -yOffset)
            toView.transform = CGAffineTransform.identity
        }, completion: { _ in
            if let viewController = viewController as? AbstractViewController {
                viewController.parallaxView.transform = CGAffineTransform(translationX: -xOffset, y: -yOffset)
                UIView.animate(withDuration: animationDuration, delay: 0, options: .curveEaseInOut, animations: {
                    viewController.parallaxView.transform = CGAffineTransform.identity
                })
            }
        })
    }
}
